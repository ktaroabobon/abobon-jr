package services

import (
	"app/internal/models"
	"app/internal/repositories"
	"app/internal/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type ThesisService struct {
	Repo   *repositories.CiNiiRepository
	Logger *utils.Logger
}

func NewThesisService(logger *utils.Logger) *ThesisService {
	return &ThesisService{Repo: repositories.NewCiNiiRepository(os.Getenv("CINII_APP_ID"), logger), Logger: logger}
}

// handleThesisCommand関数
func (s *ThesisService) HandleThesisCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) (reply discordgo.InteractionResponse) {
	// キーワードを取得
	keyword := interaction.ApplicationCommandData().Options[0].StringValue()

	// CiNii APIを使って論文を検索
	data, err := s.Repo.FetchThesis(keyword)
	if err != nil {
		s.Logger.ErrorLogger.Printf("FetchThesis response: %v", *data)
		s.Logger.ErrorLogger.Printf("論文検索中にエラーが発生しました: %v", err)
		return
	}

	var papers []models.Paper
	for _, item := range data.Items {
		// IDをURLから抽出
		idStr := strings.TrimPrefix(item.ID, "https://cir.nii.ac.jp/crid/")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			s.Logger.ErrorLogger.Printf("IDの解析に失敗しました: %v", err)
			return
		}

		// 著者をカンマ区切りに変換
		authors := strings.Join(item.Creators, ", ")

		// 出版日を解析
		publicationDate := time.Time{}
		if item.PublicationDate != "" {
			var layout string
			switch len(item.PublicationDate) {
			case 4: // YYYY
				layout = "2006"
			case 7: // YYYY-MM
				layout = "2006-01"
			default: // YYYY-MM-DD
				layout = "2006-01-02"
			}

			pd, err := time.Parse(layout, item.PublicationDate)
			if err != nil {
				s.Logger.ErrorLogger.Printf("出版日の解析に失敗しました: %v", err)
				return
			}
			publicationDate = pd
		}

		// NAIDを抽出
		NAID := extractIdentifier(item.Identifiers, "cir:NAID")

		// Paper構造体を作成
		paper := models.Paper{
			ID:              uint(id),
			Title:           item.Title,
			Authors:         &authors,
			PublicationDate: &publicationDate,
			Publisher:       &item.Publisher,
			PublicationName: &item.PublicationName,
			DOI:             extractIdentifier(item.Identifiers, "cir:DOI"),
			NAID:            &NAID,
			URL:             item.Link.ID,
		}

		papers = append(papers, paper)
	}

	// papersをInfoLoggerに出力
	s.Logger.InfoLogger.Printf("FetchThesis response: %v", papers)
	embeds := make([]*discordgo.MessageEmbed, 0)
	for _, paper := range papers {
		embed := &discordgo.MessageEmbed{
			Color: 0x00ff00,
			Title: paper.Title,
			URL:   paper.URL,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "公開日時",
					Value:  paper.PublicationDate.Format(time.RFC3339),
					Inline: true,
				},
				{
					Name:   "著者",
					Value:  *paper.Authors,
					Inline: true,
				},
			},
		}

		embeds = append(embeds, embed)
	}

	return discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	}
}

// extractIdentifier は特定のタイプの識別子を検索します
func extractIdentifier(identifiers []struct {
	Type  string `json:"@type"`
	Value string `json:"@value"`
}, idType string) string {
	for _, id := range identifiers {
		if id.Type == idType {
			return id.Value
		}
	}
	return ""
}

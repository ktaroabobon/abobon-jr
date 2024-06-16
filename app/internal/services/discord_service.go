package services

import (
	"app/internal/repositories"
	"app/internal/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mmcdole/gofeed"
)

// Articleの構造体を定義
type Article struct {
	Title       string
	Link        string
	Description string
	PublishedOn time.Time
}

type DiscordService struct {
	Repo   *repositories.DiscordRepository
	Logger *utils.Logger
}

func NewDiscordService(logger *utils.Logger) *DiscordService {
	// Repository層のインスタンス作成
	repo := repositories.NewDiscordRepository(&gofeed.Parser{})

	return &DiscordService{Repo: repo, Logger: logger}
}

// itemからArticle structを作成する関数
func (s *DiscordService) NewArticle(item *gofeed.Item) *Article {
	summary := s.summarizeContent(item.Description)

	return &Article{
		Title:       item.Title,
		Link:        item.Link,
		Description: summary,
		PublishedOn: s.safeParsePublishedDate(item.PublishedParsed),
	}
}

// PublishedParsedがnilの場合にtime.Time{}を返す関数
func (s *DiscordService) safeParsePublishedDate(PublishedParsed *time.Time) time.Time {
	if PublishedParsed != nil {
		return *PublishedParsed
	}
	return time.Time{}
}

// handlePingCommand関数
func (s *DiscordService) HandlePingCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}

	err := session.InteractionRespond(interaction.Interaction, response)
	if err != nil {
		s.Logger.ErrorLogger.Printf("Pingコマンドへの応答中にエラーが発生しました: %v", err)
		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Pingコマンドの処理中にエラーが発生しました: %v", err),
			},
		}
		err = session.InteractionRespond(interaction.Interaction, response)
		if err != nil {
			s.Logger.ErrorLogger.Printf("エラーメッセージの送信中にエラーが発生しました: %v", err)
		}
	}
}

// handleArticleCommand関数
func (s *DiscordService) HandleArticleCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	rssItems, err := s.Repo.FetchRSS("https://ktaroabobon.github.io/index.xml")
	if err != nil {
		s.Logger.ErrorLogger.Printf("RSSフィードの取得中にエラーが発生しました: %v", err)
		return
	}

	articles := make([]*discordgo.MessageEmbed, 0)
	for _, item := range rssItems {
		article := s.NewArticle(item)
		embed := &discordgo.MessageEmbed{
			Color:       0x00ff00,
			Title:       article.Title,
			URL:         article.Link,
			Description: article.Description,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "公開日時",
					Value:  article.PublishedOn.Format(time.RFC3339),
					Inline: true,
				},
			},
		}

		articles = append(articles, embed)
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: articles,
		},
	}

	err = session.InteractionRespond(interaction.Interaction, response)
	if err != nil {
		s.Logger.ErrorLogger.Printf("記事コマンドへの応答中にエラーが発生しました: %v", err)
		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("記事コマンドの処理中にエラーが発生しました: %v", err),
			},
		}
		err = session.InteractionRespond(interaction.Interaction, response)
		if err != nil {
			s.Logger.ErrorLogger.Printf("エラーメッセージの送信中にエラーが発生しました: %v", err)
		}
	}
}

// summarizeContent関数
func (s *DiscordService) summarizeContent(content string) string {
	const maxSummaryLength = 200
	if len(content) > maxSummaryLength {
		return content[:maxSummaryLength] + "..."
	}
	return content
}

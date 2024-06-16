package services

import (
	"app/internal/repositories"
	"app/internal/utils"
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

type AbobonArticlesService struct {
	Repo   *repositories.PaperRepository
	Logger *utils.Logger
}

func NewAbobonArticlesService(logger *utils.Logger) *AbobonArticlesService {
	return &AbobonArticlesService{
		Repo:   repositories.NewPaperRepository(gofeed.NewParser(), logger),
		Logger: logger,
	}
}

// itemからArticle structを作成する関数
func (s *AbobonArticlesService) NewArticle(item *gofeed.Item) *Article {
	summary := s.summarizeContent(item.Description)

	return &Article{
		Title:       item.Title,
		Link:        item.Link,
		Description: summary,
		PublishedOn: s.safeParsePublishedDate(item.PublishedParsed),
	}
}

// summarizeContent関数
func (s *AbobonArticlesService) summarizeContent(content string) string {
	const maxSummaryLength = 200
	if len(content) > maxSummaryLength {
		return content[:maxSummaryLength] + "..."
	}
	return content
}

// PublishedParsedがnilの場合にtime.Time{}を返す関数
func (s *AbobonArticlesService) safeParsePublishedDate(PublishedParsed *time.Time) time.Time {
	if PublishedParsed != nil {
		return *PublishedParsed
	}
	return time.Time{}
}

// handleArticleCommand関数
func (s *AbobonArticlesService) HandleAbobonArticlesCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) (reply discordgo.InteractionResponse) {
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

	return discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: articles,
		},
	}
}

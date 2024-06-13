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

type DiscordService struct {
	Repo   *repositories.DiscordRepository
	Logger *utils.Logger
}

func NewDiscordService(repo *repositories.DiscordRepository, logger *utils.Logger) *DiscordService {
	return &DiscordService{Repo: repo, Logger: logger}
}

// itemからArticle structを作成する関数
func (s *DiscordService) NewArticle(item *gofeed.Item) *Article {
	summary := s.summarizeContent(item.Description)

	return &Article{
		Title:       item.Title,
		Link:        item.Link,
		Description: summary,
		PublishedOn: *item.PublishedParsed,
	}
}

// handlePingCommand and handleArticleCommand methods go here
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
		s.Logger.ErrorLogger.Printf("Error responding to ping command: %v", err)
	}
}

// handleArticleCommand関数
func (s *DiscordService) HandleArticleCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	rssItems, err := s.Repo.FetchRSS("https://ktaroabobon.github.io/index.xml")
	if err != nil {
		s.Logger.ErrorLogger.Printf("Error fetching RSS feed: %v", err)
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
		s.Logger.ErrorLogger.Printf("Error responding to article command: %v", err)
	}
}

// summarizeContent関数
func (s *DiscordService) summarizeContent(content string) string {
	var sectionText string

	if sectionText == "" {
		const maxSummaryLength = 200
		if len(content) > maxSummaryLength {
			return content[:maxSummaryLength] + "..."
			// 200文字以上の場合は、最初の200文字までを返す
		}
	}

	return sectionText + "..."
}

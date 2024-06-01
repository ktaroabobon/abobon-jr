package controllers

import (
	"app/internal/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mmcdole/gofeed"
)

type DiscordController struct {
	Session *discordgo.Session
	Logger  *utils.Logger
}

func NewDiscordController(session *discordgo.Session, logger *utils.Logger) *DiscordController {
	return &DiscordController{Session: session, Logger: logger}
}

func (c *DiscordController) RegisterCommands() {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Replies with pong",
		},
		{
			Name:        "abobon-articles",
			Description: "Fetch the data from https://ktaroabobon.github.io/index.xml with RSS",
		},
	}

	for _, cmd := range commands {
		_, err := c.Session.ApplicationCommandCreate(c.Session.State.User.ID, "", cmd)
		if err != nil {
			c.Logger.ErrorLogger.Fatalf("Cannot create '%v' command: %v", cmd.Name, err)
		}
	}
}

func (c *DiscordController) HandleSlashCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "ping":
			c.handlePingCommand(s, i)
		case "abobon-articles":
			c.handleArticleCommand(s, i)
		}
	}
}

func (c *DiscordController) handlePingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}

	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		c.Logger.ErrorLogger.Printf("Error responding to ping command: %v", err)
	}
}

func (c *DiscordController) handleArticleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	items, err := fetchRSS("https://ktaroabobon.github.io/index.xml")
	if err != nil {
		c.Logger.ErrorLogger.Printf("Error fetching RSS feed: %v", err)
		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error fetching articles.",
			},
		}
		err = s.InteractionRespond(i.Interaction, response)
		if err != nil {
			c.Logger.ErrorLogger.Printf("Error responding to error message of abobon-articles command: %v", err)
		}
		return
	}

	// 最新の5件の記事を表示
	content := "Latest articles:\n"
	for i, item := range items {
		if i >= 5 {
			break
		}
		content += fmt.Sprintf("[%d] %s\n", i+1, item.Title)
	}

	// TODO: レスポンス部分
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	}

	err = s.InteractionRespond(i.Interaction, response)
	if err != nil {
		c.Logger.ErrorLogger.Printf("Error responding to abobon-articles command: %v", err)
	}
}

func fetchRSS(url string) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return feed.Items, nil
}

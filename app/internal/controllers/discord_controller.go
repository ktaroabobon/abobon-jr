package controllers

import (
	"app/internal/services"
	"app/internal/utils"

	"github.com/bwmarrin/discordgo"
)

type DiscordController struct {
	Session *discordgo.Session
	Service *services.DiscordService
}

func NewDiscordController(session *discordgo.Session, logger *utils.Logger) *DiscordController {
	// Service層のインスタンス作成
	service := services.NewDiscordService(logger)

	return &DiscordController{
		Session: session,
		Service: service,
	}
}

// RegisterCommands関数
func (c *DiscordController) RegisterCommands() {
	if c.Session.State == nil || c.Session.State.User == nil {
		c.Service.Logger.ErrorLogger.Fatal("Discord session state or user is nil")
		return
	}

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
			c.Service.Logger.ErrorLogger.Fatalf("Cannot create '%v' command: %v", cmd.Name, err)
		}
	}
}

// HandleSlashCommands関数
func (c *DiscordController) HandleSlashCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "ping":
			c.Service.HandlePingCommand(s, i)
		case "abobon-articles":
			c.Service.HandleArticleCommand(s, i)
		}
	}
}

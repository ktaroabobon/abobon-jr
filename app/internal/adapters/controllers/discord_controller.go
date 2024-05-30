package controllers

import (
	"app/internal/utils"
	"github.com/bwmarrin/discordgo"
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
		// switch i.ApplicationCommandData().Name {
		// case "ping":
		// 	c.handlePingCommand(s, i)
		// }
		if i.ApplicationCommandData().Name == "ping" {
			c.handlePingCommand(s, i)
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

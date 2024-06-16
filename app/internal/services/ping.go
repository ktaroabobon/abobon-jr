package services

import (
	"app/internal/utils"

	"github.com/bwmarrin/discordgo"
)

type PingService struct {
	Logger *utils.Logger
}

func NewPingService(logger *utils.Logger) *PingService {
	return &PingService{Logger: logger}
}

// handlePingCommand関数
func (s *PingService) HandlePingCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) (reply discordgo.InteractionResponse) {
	return discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}
}

package services

import (
	"app/internal/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordService struct {
	Session               *discordgo.Session
	Interaction           *discordgo.InteractionCreate
	Logger                *utils.Logger
	PingService           *PingService
	AbobonArticlesService *AbobonArticlesService
	ThesisService         *ThesisService
}

func NewDiscordService(logger *utils.Logger) *DiscordService {
	return &DiscordService{
		Logger:                logger,
		PingService:           NewPingService(logger),
		AbobonArticlesService: NewAbobonArticlesService(logger),
		ThesisService:         NewThesisService(logger),
	}
}

// HandleInteraction関数
func (s *DiscordService) HandleInteraction(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	s.Session = session
	s.Interaction = interaction

	commandName := interaction.ApplicationCommandData().Name
	var reply discordgo.InteractionResponse
	s.Logger.InfoLogger.Printf("'%s'コマンドが実行されました", commandName)
	switch interaction.ApplicationCommandData().Name {
	case "ping":
		reply = s.PingService.HandlePingCommand(session, interaction)
	case "abobon-articles":
		reply = s.AbobonArticlesService.HandleAbobonArticlesCommand(session, interaction)
	case "thesis":
		reply = s.ThesisService.HandleThesisCommand(session, interaction)
	}

	err := session.InteractionRespond(interaction.Interaction, &reply)
	if err != nil {
		s.Logger.ErrorLogger.Printf("%sコマンドへの応答中にエラーが発生しました: %v", commandName, err)
		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("%sコマンドの処理中にエラーが発生しました: %v", commandName, err),
			},
		}
		err = session.InteractionRespond(interaction.Interaction, response)
		if err != nil {
			s.Logger.ErrorLogger.Printf("エラーメッセージの送信中にエラーが発生しました: %v", err)
		}
	}
}

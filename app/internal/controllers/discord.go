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
		c.Service.Logger.ErrorLogger.Fatal("Discordのセッションの状態またはユーザーがnilです")
		return
	}

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "pongと返信します",
		},
		{
			Name:        "abobon-articles",
			Description: "https://ktaroabobon.github.io/index.xml からデータを取得します（RSS）",
		},
		{
			Name:        "thesis",
			Description: "キーワードを指定して論文を検索します",
		},
	}

	for _, cmd := range commands {
		_, err := c.Session.ApplicationCommandCreate(c.Session.State.User.ID, "", cmd)
		if err != nil {
			c.Service.Logger.ErrorLogger.Fatalf("'%v'コマンドを作成できませんでした: %v", cmd.Name, err)
		}
	}
}

// HandleSlashCommands関数
func (c *DiscordController) HandleSlashCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		c.Service.HandleInteraction(s, i)
	}
}

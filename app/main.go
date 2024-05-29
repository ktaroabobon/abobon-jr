package main

import (
	"os"
	"os/signal"
	"syscall"

	"app/internal/adapters/controllers"
	"app/internal/utils"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// ロガーの作成
	logger := utils.NewLogger()

	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		logger.ErrorLogger.Fatal("No DISCORD_BOT_TOKEN provided")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.ErrorLogger.Fatal("Error creating Discord session: ", err)
	}

	// Discordコントローラーの作成
	discordController := controllers.NewDiscordController(dg, logger)

	// スラッシュコマンドのハンドラを追加
	dg.AddHandler(discordController.HandleSlashCommands)

	// Discordセッションの開始
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Fatal("Error opening Discord session: ", err)
	}

	// スラッシュコマンドの登録
	discordController.RegisterCommands()

	logger.InfoLogger.Println("Bot is running. Press CTRL+C to exit.")

	// シグナルを待つ
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	// Discordセッションのクリーンな終了
	dg.Close()
}

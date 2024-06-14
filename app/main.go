package main

import (
	"os"
	"os/signal"
	"syscall"

	"app/internal/controllers"
	"app/internal/utils"

	"github.com/bwmarrin/discordgo"
)

func main() {
	logger, discordController, err := setup()
	if err != nil {
		logger.ErrorLogger.Fatalf("Error setting up bot: %v", err)
		return
	}

	// コマンドの登録
	discordController.RegisterCommands()

	logger.InfoLogger.Println("Bot is now running. Press CTRL+C to exit.")

	// シグナルを受け取るまで待機
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	// Discordセッションのクローズ
	discordController.Session.Close()
}

func setup() (*utils.Logger, *controllers.DiscordController, error) {
	logger := utils.NewLogger()

	// DISCORD_BOT_TOKENの取得
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		logger.ErrorLogger.Fatal("DISCORD_BOT_TOKEN is not set")
	}

	// Discordセッションの作成
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error creating Discord session: %v", err)
		return nil, nil, err
	}

	// DiscordControllerのインスタンスを作成
	discordController := controllers.NewDiscordController(dg, logger)

	// スラッシュコマンドのハンドラを登録
	dg.AddHandler(discordController.HandleSlashCommands)

	// Discordセッションの開始
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Fatalf("Error opening Discord session: %v", err)
		return nil, nil, err
	}

	return logger, discordController, nil
}

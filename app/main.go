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
		logger.ErrorLogger.Fatalf("ボットの設定中にエラーが発生しました: %v", err)
		return
	}

	// コマンドの登録
	discordController.RegisterCommands()

	logger.InfoLogger.Println("ボットが起動しました。CTRL+C を押して終了します。")

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
		logger.ErrorLogger.Fatal("DISCORD_BOT_TOKENが設定されていません。")
	}

	// Discordセッションの作成
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.ErrorLogger.Fatalf("Discordセッションの作成中にエラーが発生しました: %v", err)
		return nil, nil, err
	}

	// DiscordControllerのインスタンスを作成
	discordController := controllers.NewDiscordController(dg, logger)

	// スラッシュコマンドのハンドラを登録
	dg.AddHandler(discordController.HandleSlashCommands)

	// Discordセッションの開始
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Fatalf("Discordセッションのオープン中にエラーが発生しました: %v", err)
		return nil, nil, err
	}

	return logger, discordController, nil
}

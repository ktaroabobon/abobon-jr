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

	// configの読み込み
	config, err := utils.NewConfig()
	if err != nil {
		logger.ErrorLogger.Fatalf("configの読み込み中にエラーが発生しました: %v", err)
	}

	// Discordセッションの作成
	dg, err := discordgo.New("Bot " + config.DiscordBotToken)
	if err != nil {
		logger.ErrorLogger.Fatalf("Discordセッションの作成中にエラーが発生しました: %v", err)
	}

	// DiscordControllerのインスタンスを作成
	discordController := controllers.NewDiscordController(dg, config, logger)

	// スラッシュコマンドのハンドラを登録
	dg.AddHandler(discordController.HandleSlashCommands)

	// Discordセッションの開始
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Fatalf("Discordセッションのオープン中にエラーが発生しました: %v", err)
	}

	return logger, discordController, nil
}

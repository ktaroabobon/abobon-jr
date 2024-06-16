package utils

import (
	"fmt"
	"os"
)

type Config struct {
	DiscordBotToken string
	DiscordClientID string
	DiscordGuildID  string
	CiniiAppID      string
}

func GetEnvVar(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("環境変数 %s が存在しません", key)
	}
	if value == "" {
		return "", fmt.Errorf("環境変数 %s の値が空です", key)
	}
	return value, nil
}

func NewConfig() (*Config, error) {
	// discord bot tokenの取得
	token, err := GetEnvVar("DISCORD_BOT_TOKEN")
	if err != nil {
		return nil, err
	}

	// discord client idの取得
	clientID, err := GetEnvVar("DISCORD_CLIENT_ID")
	if err != nil {
		return nil, err
	}

	// discord guild idの取得
	guildID, err := GetEnvVar("DISCORD_GUILD_ID")
	if err != nil {
		return nil, err
	}

	// appIDの取得
	appID, err := GetEnvVar("CINII_APP_ID")
	if err != nil {
		return nil, err
	}

	return &Config{
		DiscordBotToken: token,
		DiscordClientID: clientID,
		DiscordGuildID:  guildID,
		CiniiAppID:      appID,
	}, nil
}

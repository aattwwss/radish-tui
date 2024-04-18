package main

import (
	"fmt"

	"github.com/aattwwss/radish-tui/reddit"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	// reddit config
	ClientId     string `env:"CLIENT_ID,notEmpty"`
	ClientSecret string `env:"CLIENT_SECRET,notEmpty"`
	Username     string `env:"BOT_USERNAME,notEmpty"`
	Password     string `env:"BOT_PASSWORD,notEmpty"`

	//debugging config
	Token           string `env:"BOT_ACCESS_TOKEN"`
	ExpireTimeMilli int64  `env:"BOT_TOKEN_EXPIRE_MILLI"`
	IsDebug         bool   `env:"IS_DEBUG"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Msgf("Parse env error: %v", err)
	}

	rc, err := reddit.NewRedditClient(cfg.ClientId, cfg.ClientSecret, cfg.Username, cfg.Password, cfg.Token, cfg.ExpireTimeMilli)
	if err != nil {
		log.Fatal().Msgf("Init reddit client error: %v", err)
	}
	submissions, err := rc.GetSubmissions("pcgaming", reddit.HOT, 10)
	if err != nil {
		log.Error().Msgf("error: %v", err)
	}
	for _, submission := range submissions {
		fmt.Println(submission.Title)
	}

}

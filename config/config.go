package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
)

type appconfig struct {
	Name string `yaml:"name" env:"NAME"`
	Port string `yaml:"port" env:"PORT"`
	Line struct {
		ChannelSecret string `yaml:"channel_secret" env:"CHANNEL_SECRET"`
		ChannelToken  string `yaml:"channel_token" env:"CHANNEL_TOKEN"`
	} `yaml:"line"`
}

var App *appconfig

func init() {
	App = &appconfig{}
}

func Init(path string) {
	err := cleanenv.ReadConfig(path+"app.yaml", App)
	if err != nil {
		log.Fatal().Err(err).Msg("Application configuration initialize failed, app.yaml has a problem")
		return
	}
	_ = cleanenv.ReadEnv(App)
}

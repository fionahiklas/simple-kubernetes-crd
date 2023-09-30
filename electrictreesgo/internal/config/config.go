package config

import (
	"github.com/caarlos0/env/v6"
)

type treesConfig struct {
	TreeName   string `env:"TREE_NAME" envDefault:"Sparkles"`
	Try        bool   `env:"TREE_TRY" envDefault:"false"`
	HowFarAway int    `env:"TREE_DISTANCE" envDefault:"0"`
	EyesClosed bool   `env:"TREE_EYES_CLOSED" envDefault:"false"`
}

type loggingConfig struct {
	LogLevel            string `env:"LOG_LEVEL" envDefault:"INFO"`
	EnableJsonLogFormat bool   `env:"ENABLE_JSON_LOG_FORMAT" envDefault:"false"`
}

type config struct {
	Logging loggingConfig
	Tree    treesConfig
}

func ParseConfig() (*config, error) {
	config := &config{}
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (c *config) LogLevel() string {
	return c.Logging.LogLevel
}

func (c *config) EnableJsonLogFormat() bool {
	return c.Logging.EnableJsonLogFormat
}

func (c *config) TreeName() string {
	return c.Tree.TreeName
}

func (c *config) Try() bool {
	return c.Tree.Try
}

func (c *config) HowFarAway() int {
	return c.Tree.HowFarAway
}

func (c *config) EyesClosed() bool {
	return c.Tree.EyesClosed
}

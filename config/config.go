package config

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel zapcore.Level `env:"LOG_LEVEL" envDefault:"debug"`
	JsonLogs bool          `env:"JSON_LOGS" envDefault:"false"`
	OneShot  bool          `env:"ONE_SHOT" envDefault:"false"`

	DatabaseUrl string `env:"DATABASE_URL,required"`
}

var Conf Config

func Parse() {
	if err := env.Parse(&Conf); err != nil {
		panic(err)
	}
}

package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost         string `mapStructure:"POSTGRES_HOST"`
	DBUserName     string `mapStructure:"POSTGRES_USER"`
	DBUserPassword string `mapStructure:"POSTGRES_PASSWORD"`
	DBName         string `mapStructure:"POSTGRES_DB"`
	DBPort         string `mapStructure:"POSTGRES_PORT"`

	JwtSecret    string        `mapStructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapStructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapStructure:"JWT_MAXAGE"`

	ClientOrigin string `mapStructure:"CLIENT_ORIGIN"`
}

func LoadEnv(path string) (Env Env, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Env)
	return
}

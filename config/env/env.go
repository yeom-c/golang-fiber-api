package env

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type env struct {
	ServerPort int `mapstructure:"SERVER_PORT"`
}

var Env env

func setDefault() {
	viper.SetDefault("SERVER_PORT", 5000)
}

func InitEnv(path string) {
	serverEvn := os.Getenv("SERVER_ENV")
	if serverEvn == "" {
		serverEvn = "local"
	}

	viper.SetConfigName(serverEvn)
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	// 동일한 이름의 변수는 시스템 환경변수로 덮어씀.
	viper.AutomaticEnv()
	// 데이터 설정이 안된 변수들 기본값 설정.
	setDefault()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to load env file. %+v", err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		log.Fatalf("Failed to unmarshal env. %+v", err)
	}
}

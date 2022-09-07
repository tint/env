package env

import (
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"strings"
)

var (
	RootPath string
	AppEnv   string
)

func Setup(root ...string) {
	if len(root) == 0 {
		path, err := filepath.Abs(".")
		if err != nil {
			log.Fatal(err)
		}
		RootPath = path
	} else {
		RootPath, _ = filepath.Abs(root[0])
	}

	// 加载环境变量 .env .env.{APP_ENV}
	godotenv.Load(RootPath + "/.env")
	AppEnv = String("APP_ENV", "release")
	if len(AppEnv) > 0 {
		currEnvFile := ".env." + strings.ToLower(AppEnv)
		godotenv.Overload(RootPath + "/" + currEnvFile)
	}
}

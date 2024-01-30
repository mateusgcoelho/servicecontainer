package app

import "github.com/joho/godotenv"

func InitConfigEnvs() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

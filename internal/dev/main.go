package main

import (
	"github.com/melkor217/bitmagnet/internal/dev/app"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.New().Run()
}

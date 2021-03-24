package main

import (
	"github.com/joho/godotenv"
	conf "github.com/yuriimakohon/go-chat/config"
	"github.com/yuriimakohon/go-chat/internal/handler"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"github.com/yuriimakohon/go-chat/internal/repository/psql"
	"github.com/yuriimakohon/go-chat/internal/server"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(conf.EnvPath); err != nil {
		log.Fatal("Config file wasn't load")
	}

	var repo repository.Repository
	repo = psql.New()
	if repo == nil {
		log.Fatal("Repository was not created")
	}

	h := handler.New(repo)
	s := new(server.Server)
	err := s.Run(
		os.Getenv("PORT"),
		h.SetupRoutes())
	if err != nil {
		log.Fatalf("error occurred while running program http server: %s", err)
	}
}

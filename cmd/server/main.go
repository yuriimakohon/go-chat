package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	conf "github.com/yuriimakohon/go-chat/configs"
	"github.com/yuriimakohon/go-chat/internal/handler"
	"github.com/yuriimakohon/go-chat/internal/repository"
	"github.com/yuriimakohon/go-chat/internal/repository/postgres"
	"github.com/yuriimakohon/go-chat/internal/server"
	"github.com/yuriimakohon/go-chat/internal/service"
	"log"
)

func main() {
	if err := intiConfigs(); err != nil {
		log.Fatal("config files wasn't load: ", err.Error())
	}
	if err := godotenv.Load(conf.EnvPath); err != nil {
		log.Fatal("environments vars file wasn't load: ", err.Error())
	}

	db, err := postgres.NewDB()
	if err != nil {
		log.Fatal("error occurred while creating db: ", err.Error())
	}

	chatRepository := repository.NewRepository(
		postgres.NewAuthRepository(db),
		postgres.NewRoomRepository(db))
	chatService := service.NewService(chatRepository)
	chatHandler := handler.NewHandler(chatService)

	serv := new(server.Server)
	if err := serv.Run(viper.GetString("port"), chatHandler.InitRoutes()); err != nil {
		log.Fatal("error occurred while running server: ", err.Error())
	}
}

func intiConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

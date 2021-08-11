package main

import (
	"github.com/HunkevychPhilip/todo/pkg/handler"
	"github.com/HunkevychPhilip/todo/pkg/repository"
	"github.com/HunkevychPhilip/todo/pkg/service"
	"github.com/HunkevychPhilip/todo/pkg/utils"
	"github.com/HunkevychPhilip/todo/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Could not read config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Cound not read env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Cannot establish db conection: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	errHandler := utils.NewErrHandler()
	utilities := utils.NewUtils(errHandler)
	handlers := handler.NewHandler(services, utilities)

	s := new(server.Server)
	if err := s.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Server returned an error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

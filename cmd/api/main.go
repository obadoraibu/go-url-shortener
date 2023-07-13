package main

import (
	"flag"
	"github.com/obadoraibu/go-url-shortener/internal/repository"
	"github.com/obadoraibu/go-url-shortener/internal/service"
	"github.com/obadoraibu/go-url-shortener/internal/transport/rest"
	"github.com/obadoraibu/go-url-shortener/internal/transport/rest/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	flag.Parse()

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error reading configs file, %s", err)
	}

	repo, err := repository.NewRepository()
	if err != nil {
		logrus.Fatalf("error creating repository, %s", err)
	}

	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	server := rest.NewServer()

	if err := server.Start(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	//go func() {
	//	if err := server.Start(viper.GetString("port"), handler.InitRoutes()); err != nil {
	//		logrus.Fatalf("error occured while running http server: %s", err.Error())
	//	}
	//}()
	//
	//logrus.Print("url-shortener started")
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//<-quit
	//
	//logrus.Print("url shutting down")
	//
	//if err := server.Stop(context.Background()); err != nil {
	//	logrus.Errorf("error on server shutting down: %s", err.Error())
	//}
	//
	//if err := repo.Close(); err != nil {
	//	logrus.Errorf("error on db connection close: %s", err.Error())
	//}
}

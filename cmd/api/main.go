package main

import (
	"telefun/config"
	"telefun/handler"
	"telefun/repository"
	"telefun/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Connect(cfg.DB.Driver, cfg.DB.DSN())
	if err != nil {
		panic(err)
	}

	repo := repository.NewPhoneNumberRepository(db)
	usecase := usecase.NewPhoneNumberUsecase(repo)
	handler := handler.NewPhoneNumberHandler(usecase)

	r := gin.Default()
	r.POST("/v1/phone_numbers", handler.Create)
	r.GET("/v1/phone_numbers/available", handler.Available)

	if err := r.Run(cfg.HttpServer.Addr); err != nil {
		panic(err)
	}
}

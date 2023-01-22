package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() {

	dsn := "host=localhost user=postgres password=postgres dbname=demo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	userRepo := NewReposiory(db)
	userService := NewUserService(userRepo)

	r := gin.New()
	r.Use(gin.Recovery())

	// register handler
	RegisterUserHTTPHandler(r, userService)

	r.Run()

}

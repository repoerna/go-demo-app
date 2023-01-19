package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Demo Apps")

	dsn := "host=localhost user=postgres password=postgres dbname=demo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(User{})

	repo := NewReposiory(db)

	uuid, err := uuid.Parse("3e88461a-5cf3-4f6c-937f-478c45a4cc55")
	if err != nil {
		log.Fatal(err)
	}

	user, err := repo.GetUser(uuid)
	fmt.Println(err)
	fmt.Println(user)
}

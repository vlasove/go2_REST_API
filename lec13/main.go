package main

import (
	"lec13/config"
	"lec13/models"
	"lec13/routes"

	"github.com/jinzhu/gorm"

	"fmt"
)

var err error //глобальный отлов ошибок

func main() {
	config.DB, err = gorm.Open("postgres", config.DbURI(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Current status of app:", err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Run(":8080")
}

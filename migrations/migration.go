package main

import (
	"gin-api/infra"
	"gin-api/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("failed to migrate database")
	}
}

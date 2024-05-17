package main

import (
	"gin-api/infra"
	"gin-api/repositories"

	"gin-api/controllers"
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// items := []models.Item{
	// 	{ID: 1, Name: "item1", Price: 100, Description: "This is item1", SoldOut: false},
	// 	{ID: 2, Name: "item2", Price: 200, Description: "This is item2", SoldOut: true},
	// 	{ID: 3, Name: "item3", Price: 300, Description: "This is item3", SoldOut: false},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("items", itemController.Create)
	r.PUT("items/:id", itemController.Update)
	r.DELETE("items/:id", itemController.Delete)
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}

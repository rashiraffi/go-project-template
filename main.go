package main

import (
	"log"
	"net/http"
	"template/database"
	"template/handler"
	"template/model"
	"template/router"
	"template/service"
)

func main() {
	db := database.Connect()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	model := model.NewModel(db)
	service := service.NewService(model)
	handler := handler.NewHandler(service)

	router := router.GetRouter(handler)
	log.Default().Println("Server started on port 8080")
	http.ListenAndServe(":8080", router)
}

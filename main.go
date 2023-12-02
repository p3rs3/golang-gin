package main

import (
	controller "crud-gin/controllers"
	repository "crud-gin/repositories"
	"crud-gin/routes"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title 	Users Service API
// @version	1.0
// @description A Users service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /docs
func main() {
	dsn := "host=176.57.217.75 user=gen_user password={:0TJX5Qsyt~GB dbname=default_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	repository := repository.CreateUsersRepository(db)
	controller := controller.CreateUsersController(repository)
	routes := routes.CreateUsersRoutes(controller)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes.GetRoutes(),
	}
	server.ListenAndServe()
}

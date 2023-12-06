package main

import (
	"crud-gin/config"
	controller "crud-gin/controllers"
	repository "crud-gin/repositories"
	"crud-gin/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

// @title 		Users Service API
// @version		1.0
// @description A Users service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /docs
func main() {
	host, _ := os.LookupEnv("DB_HOST")

	fmt.Println(host)

	config := config.GetConfig()

	dsn := "host=" + config.Db.Host +
		" user=" + config.Db.User +
		" password=" + config.Db.Password +
		" dbname=" + config.Db.Name +
		" port=" + config.Db.Port +
		" sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	repository := repository.CreateUsersRepository(db)
	controller := controller.CreateUsersController(repository)
	routes := routes.CreateUsersRoutes(controller)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes.GetRoutes(),
	}
	server.ListenAndServe()
}

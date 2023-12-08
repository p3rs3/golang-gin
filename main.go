package main

import (
	"crud-gin/config"
	controller "crud-gin/controllers"
	repository "crud-gin/repositories"
	"crud-gin/router"
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

// @title 		API
// @version		1.0

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

	router := router.CreateRouter()

	routes.CreateUsersRoutes(controller).SetRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

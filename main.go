package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/kwamekyeimonies/ecommerce_web_app/controllers"
	// "github.com/kwamekyeimonies/ecommerce_web_app/database"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductsData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	log.Fatal(router.Run(":" + port))
}

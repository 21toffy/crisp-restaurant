package main

import (
	"os"

	"github.com/21toffy/crisp-restaurant/database"
	"github.com/21toffy/crisp-restaurant/middleware"
	"github.com/21toffy/crisp-restaurant/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	// router.User(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)
}

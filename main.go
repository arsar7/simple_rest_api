package main

import (
	"example/simple_rest_api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", models.Index)
	route.GET("/users", models.GetUsers)
	route.GET("/users/:id", models.GetUser)
	route.POST("/users", models.AddUser)
	route.PATCH("/users/delete/:id", models.DeleteUser)
	route.Run("localhost:5050")
}

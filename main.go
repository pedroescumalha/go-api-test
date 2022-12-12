package main

import (
	"api/example/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/fact", controllers.GetFact)
	router.GET("/afraid/:subject", controllers.GetIsChuckNorrisAfraid)

	router.Run("localhost:9090")
}

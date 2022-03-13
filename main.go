package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func main() {
	/**
	@description Setup Server
	*/
	router := SetupRouter()

	router.GET("/api", getHelloWorld)

	router.Run("localhost:4444")
}

func SetupRouter() *gin.Engine {
	/**
	@description Init Router
	*/
	router := gin.Default()

	return router
}

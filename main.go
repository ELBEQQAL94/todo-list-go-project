package main

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	// util "github.com/restuwahyu13/gin-rest-api/utils"
)

func getHelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello, World, test, main!"})
}

func main() {
	/**
	@description Setup Server
	*/
	router := SetupRouter()

	router.GET("/api", getHelloWorld)

	// log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))
	router.Run(":4444")
}

func SetupRouter() *gin.Engine {
	/**
	@description Init Router
	*/
	router := gin.Default()

	return router
}

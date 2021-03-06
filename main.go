package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

func getHelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"say": "Hello, World!"})
}

func main() {
	/**
	@description Setup Server
	*/
	router := SetupRouter()

	router.GET("/api", getHelloWorld)

	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	/**
	@description Init Router
	*/
	router := gin.Default()

	return router
}

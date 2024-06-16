package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	vacancies := []String{"Golang: Balance", "Ruby: Basekitchen"}

	router.Run(":8080")
}

// func homeHandler() {

// }

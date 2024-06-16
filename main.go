package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	vacancies := []string{"Golang: Balance", "Ruby: Basekitchen"}

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Vacancies": vacancies,
		})
	})

	router.Run(":8080")
}

// func homeHandler() {

// }

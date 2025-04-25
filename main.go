package main

import (
	"log"
	"net/http"

	"baseFork/handlers"
	"baseFork/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := models.InitDB("./data/jobs.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	// Redirect root to /home
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	router.GET("/create-job", handlers.HomeHandler)
	router.POST("/create-job", handlers.CreateJobHandler)
	router.GET("/delete/:id", handlers.DeleteJobHandler)

	router.GET("/home", func(c *gin.Context) {
		jobs, err := models.GetAllJobs(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get jobs"})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Jobs": jobs,
		})
	})

	router.Run(":8080")
}

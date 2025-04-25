package handlers

import (
	"net/http"
	"strconv"

	"baseFork/models"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{})
}

func CreateJobHandler(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	salary, err := strconv.Atoi(c.PostForm("salary"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid salary"})
		return
	}

	db, err := models.InitDB("./data/jobs.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	err = models.InsertJob(db, title, description, salary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert job"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/home")
}

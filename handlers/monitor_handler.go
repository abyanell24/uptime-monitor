package handlers

import (
	"uptime-monitor/db"
	"uptime-monitor/models"

	"github.com/gin-gonic/gin"
)

func GetMonitors(c *gin.Context) {

	rows, err := db.DB.Query("SELECT id, url FROM monitors")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var monitors []models.Monitor

	for rows.Next() {

		var m models.Monitor

		rows.Scan(&m.ID, &m.URL)

		monitors = append(monitors, m)
	}

	c.JSON(200, monitors)
}

func CreateMonitor(c *gin.Context) {

	var m models.Monitor

	if err := c.BindJSON(&m); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	_, err := db.DB.Exec(
		"INSERT INTO monitors (url) VALUES ($1)",
		m.URL,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "monitor created"})
}
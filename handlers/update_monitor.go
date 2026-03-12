package handlers

import (
	"uptime-monitor/db"

	"github.com/gin-gonic/gin"
)

type UpdateMonitorRequest struct {
	URL string `json:"url"`
}

func UpdateMonitor(c *gin.Context) {

	id := c.Param("id")

	var req UpdateMonitorRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	_, err := db.DB.Exec(
		"UPDATE monitors SET url=$1 WHERE id=$2",
		req.URL,
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "monitor updated",
	})
}
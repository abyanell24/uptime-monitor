package handlers

import (
	"uptime-monitor/db"

	"github.com/gin-gonic/gin"
)

func DeleteMonitor(c *gin.Context) {

	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM monitors WHERE id=$1", id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "monitor deleted",
	})
}
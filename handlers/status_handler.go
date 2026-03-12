package handlers

import (
	"uptime-monitor/db"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {

	monitorID := c.Param("monitor_id")

	var status int

	err := db.DB.QueryRow(
		"SELECT status FROM checks WHERE monitor_id=$1 ORDER BY checked_at DESC LIMIT 1",
		monitorID,
	).Scan(&status)

	if err != nil {
		c.JSON(200, gin.H{
			"status": "unknown",
		})
		return
	}

	if status == 200 {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "DOWN",
		})
	}
}
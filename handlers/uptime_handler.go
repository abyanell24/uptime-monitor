package handlers

import (
	"uptime-monitor/db"

	"github.com/gin-gonic/gin"
)

func GetUptime(c *gin.Context) {

	monitorID := c.Param("monitor_id")

	var total int
	var success int

	db.DB.QueryRow(
		"SELECT COUNT(*) FROM checks WHERE monitor_id=$1",
		monitorID,
	).Scan(&total)

	db.DB.QueryRow(
		"SELECT COUNT(*) FROM checks WHERE monitor_id=$1 AND status=200",
		monitorID,
	).Scan(&success)

	if total == 0 {
		c.JSON(200, gin.H{"uptime": 0})
		return
	}

	uptime := float64(success) / float64(total) * 100

	c.JSON(200, gin.H{
		"uptime": uptime,
	})
}
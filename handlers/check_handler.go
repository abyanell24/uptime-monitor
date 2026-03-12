package handlers

import (
	"uptime-monitor/db"

	"github.com/gin-gonic/gin"
)

func SomeHandler() {
    db.DB.Query("SELECT ...")
}

func GetChecks(c *gin.Context) {

	monitorID := c.Param("monitor_id")

	rows, err := db.DB.Query(
		"SELECT status,response_time,checked_at FROM checks WHERE monitor_id=$1 ORDER BY checked_at DESC LIMIT 20",
		monitorID,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var checks []map[string]interface{}

	for rows.Next() {

		var status int
		var responseTime int
		var checkedAt string

		rows.Scan(&status, &responseTime, &checkedAt)

		checks = append(checks, gin.H{
			"status": status,
			"response_time": responseTime,
			"checked_at": checkedAt,
		})
	}

	c.JSON(200, checks)
}

func GetMonitorsHandler() { // ubah namanya biar unik
    rows, err := db.DB.Query("SELECT * FROM monitors")
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    // ...
}
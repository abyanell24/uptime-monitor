package handlers

import (
	"uptime-monitor/db"
	"time"
	
	"github.com/gin-gonic/gin"
)

type Check struct {
	ID           int       `json:"id"`
	MonitorID    int       `json:"monitor_id"`
	Status       string    `json:"status"`
	ResponseTime int64     `json:"response_time"`
	CheckedAt    time.Time `json:"checked_at"`
}

// GetChecks returns last 20 checks for a monitor
func GetChecks(c *gin.Context) {
	monitorID := c.Param("monitor_id")
	rows, err := db.DB.Query("SELECT id, monitor_id, status, response_time, checked_at FROM checks WHERE monitor_id=$1 ORDER BY checked_at DESC LIMIT 20", monitorID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var checks []Check
	for rows.Next() {
		var ch Check
		rows.Scan(&ch.ID, &ch.MonitorID, &ch.Status, &ch.ResponseTime, &ch.CheckedAt)
		checks = append(checks, ch)
	}

	c.JSON(200, checks)
}

// GetUptime calculates uptime percentage
func GetUptime(c *gin.Context) {
	monitorID := c.Param("monitor_id")

	var total, up int
	db.DB.QueryRow("SELECT COUNT(*), SUM(CASE WHEN status='UP' THEN 1 ELSE 0 END) FROM checks WHERE monitor_id=$1", monitorID).Scan(&total, &up)

	uptime := 0.0
	if total > 0 {
		uptime = (float64(up) / float64(total)) * 100
	}

	c.JSON(200, gin.H{"uptime": uptime})
}

// GetStatus returns latest status
func GetStatus(c *gin.Context) {
	monitorID := c.Param("monitor_id")

	var status string
	db.DB.QueryRow("SELECT status FROM checks WHERE monitor_id=$1 ORDER BY checked_at DESC LIMIT 1", monitorID).Scan(&status)
	if status == "" {
		status = "UNKNOWN"
	}

	c.JSON(200, gin.H{"status": status})
}
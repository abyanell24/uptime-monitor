package handlers

import (
	"net/http"
	"strconv"
	"uptime-monitor/db"

	"github.com/gin-gonic/gin"
)

type Monitor struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func GetMonitors(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, url FROM monitors")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var monitors []Monitor
	for rows.Next() {
		var m Monitor
		rows.Scan(&m.ID, &m.URL)
		monitors = append(monitors, m)
	}

	c.JSON(200, monitors)
}

func CreateMonitor(c *gin.Context) {
	var m Monitor
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.DB.QueryRow("INSERT INTO monitors (url) VALUES ($1) RETURNING id", m.URL).Scan(&m.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, m)
}

func UpdateMonitor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m Monitor
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("UPDATE monitors SET url=$1 WHERE id=$2", m.URL, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	m.ID = id
	c.JSON(200, m)
}

func DeleteMonitor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := db.DB.Exec("DELETE FROM monitors WHERE id=$1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
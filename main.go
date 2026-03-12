package main

import (
	"fmt"
	"uptime-monitor/db"
	"uptime-monitor/handlers"
	"uptime-monitor/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()
	fmt.Println("App started")
	go services.CheckWebsites() // goroutine background checker

	r := gin.Default()

	// Load templates
	r.LoadHTMLGlob("templates/*")

	// Serve static dashboard
	r.Static("/dashboard", "./static")

	// Routes: monitor CRUD
	r.GET("/monitors", handlers.GetMonitors)
	r.POST("/monitors", handlers.CreateMonitor)
	r.PUT("/monitors/:id", handlers.UpdateMonitor)
	r.DELETE("/monitors/:id", handlers.DeleteMonitor)

	// Routes: checks
	r.GET("/checks/:monitor_id", handlers.GetChecks)
	r.GET("/uptime/:monitor_id", handlers.GetUptime)
	r.GET("/status/:monitor_id", handlers.GetStatus)

	// Public status page
	r.GET("/status", func(c *gin.Context) {
		c.HTML(200, "status.html", nil)
	})

	r.Run(":8080")
}
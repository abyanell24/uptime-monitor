package services

import (
    "net/http"
    "uptime-monitor/db"
    "time"
)

func CheckWebsites() {
    for {
        rows, _ := db.DB.Query("SELECT id, url FROM monitors")
        for rows.Next() {
            var id int
            var url string
            rows.Scan(&id, &url)

            start := time.Now()
            resp, err := http.Get(url)
            duration := time.Since(start).Milliseconds()

            status := "DOWN"
            if err == nil && resp.StatusCode < 400 {
                status = "UP"
            }

            db.DB.Exec(
                "INSERT INTO checks (monitor_id, status, response_time, checked_at) VALUES ($1,$2,$3,$4)",
                id, status, duration, time.Now(),
            )
        }
        rows.Close()
        time.Sleep(10 * time.Second)
    }
}
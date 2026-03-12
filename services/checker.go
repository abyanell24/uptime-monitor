package services

import (
	"fmt"
    "net/http"
	"uptime-monitor/db"
	"time"
)

func CheckWebsites() {
	// Jangan jalan kalau DB belum siap
	if db.DB == nil {
		fmt.Println("DB is nil, skipping checks")
		return
	}

	for {
		rows, err := db.DB.Query("SELECT id, url FROM monitors")
		if err != nil {
			fmt.Println("Error fetching monitors:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		type Monitor struct {
			ID  int
			URL string
		}

		var monitors []Monitor
		for rows.Next() {
			var m Monitor
			if err := rows.Scan(&m.ID, &m.URL); err != nil {
				fmt.Println("Error scanning monitor:", err)
				continue
			}
			monitors = append(monitors, m)
		}
		rows.Close()

		// Loop setiap monitor
		for _, m := range monitors {
			start := time.Now()
			status := "DOWN"

			// Ping website
			resp, err := http.Get(m.URL)
			if err == nil && resp.StatusCode < 400 {
				status = "UP"
			}
			if resp != nil {
				resp.Body.Close()
			}

			responseTime := time.Since(start).Milliseconds()

			// Insert check
			_, err = db.DB.Exec(
				"INSERT INTO checks (monitor_id, status, response_time) VALUES ($1, $2, $3)",
				m.ID, status, responseTime,
			)
			if err != nil {
				fmt.Println("Error inserting check:", err)
			}
		}

		time.Sleep(10 * time.Second) // interval cek
	}
}
package services

import (
	"fmt"
	"net/http"
	"time"
	"uptime-monitor/db"
)

func CheckWebsites() {

	for {

		rows, err := db.DB.Query("SELECT id, url FROM monitors")

		if err != nil {
			fmt.Println("error fetching monitors:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		for rows.Next() {

			var id int
			var url string

			rows.Scan(&id, &url)

			start := time.Now()

			resp, err := http.Get(url)

			duration := time.Since(start)

			if err != nil {

				fmt.Println("DOWN:", url)

				db.DB.Exec(
					"INSERT INTO checks (monitor_id,status,response_time) VALUES ($1,$2,$3)",
					id,
					0,
					0,
				)

				continue
			}

			resp.Body.Close()

			fmt.Println("UP:", url, duration.Milliseconds(), "ms")

			db.DB.Exec(
				"INSERT INTO checks (monitor_id,status,response_time) VALUES ($1,$2,$3)",
				id,
				resp.StatusCode,
				duration.Milliseconds(),
			)
		}

		rows.Close()

		time.Sleep(30 * time.Second)
	}
}
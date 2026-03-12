package handlers

import (
    "encoding/json"
    "net/http"
    "uptime-monitor/db"
)

type Monitor struct {
    ID  int    `json:"id"`
    URL string `json:"url"`
}

func GetMonitors(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, url FROM monitors")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var monitors []Monitor
    for rows.Next() {
        var m Monitor
        rows.Scan(&m.ID, &m.URL)
        monitors = append(monitors, m)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(monitors)
}

// Tambahkan CreateMonitor, UpdateMonitor, DeleteMonitor sesuai kebutuhan
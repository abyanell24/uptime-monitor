package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "uptime-monitor/db"
)

// Monitor struct
type Monitor struct {
    ID  int    `json:"id"`
    URL string `json:"url"`
}

// GetMonitors returns all monitors
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

// CreateMonitor adds a new monitor
func CreateMonitor(w http.ResponseWriter, r *http.Request) {
    var m Monitor
    json.NewDecoder(r.Body).Decode(&m)

    err := db.DB.QueryRow("INSERT INTO monitors (url) VALUES ($1) RETURNING id", m.URL).Scan(&m.ID)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)
}

// UpdateMonitor updates the URL of an existing monitor
func UpdateMonitor(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid monitor ID", 400)
        return
    }

    var m Monitor
    json.NewDecoder(r.Body).Decode(&m)

    _, err = db.DB.Exec("UPDATE monitors SET url=$1 WHERE id=$2", m.URL, id)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    m.ID = id
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)
}

// DeleteMonitor removes a monitor
func DeleteMonitor(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid monitor ID", 400)
        return
    }

    _, err = db.DB.Exec("DELETE FROM monitors WHERE id=$1", id)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
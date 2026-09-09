package handler

import (
    "encoding/json"
    "net/http"
)

func UpsertDepartures(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"departures_updated"})
}
func UpdateDeparture(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"departure_updated"})
}
func DeleteDeparture(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"departure_deleted"})
}

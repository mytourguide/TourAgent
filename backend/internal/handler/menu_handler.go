package handler

import (
    "encoding/json"
    "net/http"
)

func ListMenus(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]interface{}{"data":[]})
}
func CreateMenu(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"created"})
}

package handler

import (
    "encoding/json"
    "net/http"
)

func ListCategories(w http.ResponseWriter, r *http.Request) {
    // Basit stub — gerçekte DB'den gelir
    cats := []map[string]string{
        {"id":"1","name":"Yurt İçi","slug":"yurt-ici"},
        {"id":"2","name":"Yurt Dışı","slug":"yurt-disi"},
    }
    json.NewEncoder(w).Encode(map[string]interface{}{"data": cats})
}

package handler

import (
    "encoding/json"
    "net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]interface{}{
        "totalBookings": 128,
        "pendingPayments": 12,
        "monthlyRevenue": 245000,
        "occupancy": 68.5,
        "dailySales": []map[string]interface{}{{"date":"2026-09-01","amount":12000},{"date":"2026-09-02","amount":18000}},
    })
}

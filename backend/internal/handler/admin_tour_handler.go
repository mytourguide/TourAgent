package handler

import (
    "encoding/json"
    "net/http"
)

type TourCreateReq struct {
    Title string `json:"title" validate:"required,max=200"`
    Slug string `json:"slug" validate:"required"`
    Description string `json:"description"`
    CategoryID string `json:"category_id"`
    BasePrice float64 `json:"base_price" validate:"gte=0"`
    Location string `json:"location"`
}

func AdminCreateTour(w http.ResponseWriter, r *http.Request) {
    var req TourCreateReq
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest); return
    }
    // validate.Struct(req) // validator ile doğrula
    json.NewEncoder(w).Encode(map[string]string{"status":"created","slug":req.Slug})
}

package handler

import (
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/repository"
)

type TourItemsPayload struct {
    Included []string `json:"included"`
    Excluded []string `json:"excluded"`
    Bring []string `json:"bring"`
    Itinerary []struct {
        DayNo int `json:"day_no"`
        Title string `json:"title"`
        Description string `json:"description"`
    } `json:"itinerary"`
    Departures []struct {
        StartDate string `json:"start_date"`
        EndDate string `json:"end_date"`
        Capacity int `json:"capacity"`
        PricePerPerson float64 `json:"price_per_person"`
    } `json:"departures"`
}

func UpsertTourItems(repo *repository.TourItemsRepo) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var p TourItemsPayload
        if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
            http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
            return
        }
        tourID := chi.URLParam(r, "id")
        // TODO: payload'ı repo formatına dönüştür
        // repo.UpsertAll(...)
        json.NewEncoder(w).Encode(map[string]string{"status":"items_updated"})
    }
}

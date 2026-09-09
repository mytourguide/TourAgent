package handler

import (
    "encoding/json"
    "net/http"
    "github.com/google/uuid"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/repository"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/models"
)

type BookingHandler struct { Repo *repository.BookingRepo }

func (h *BookingHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req struct {
        TourID string `json:"tour_id"`
        DepartureID string `json:"departure_id"`
        AdultCount int `json:"adult_count"`
        ChildCount int `json:"child_count"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, `{"error":"invalid"}`, 400); return }
    bookingNo := "TA-" + uuid.New().String()[:8]
    b := models.Booking{
        ID: uuid.NewString(),
        BookingNo: bookingNo,
        TourID: req.TourID,
        DepartureID: req.DepartureID,
        AdultCount: req.AdultCount,
        ChildCount: req.ChildCount,
        TotalPrice: 0,
        Currency: "TRY",
        Status: "pending",
    }
    if err := h.Repo.Create(r.Context(), b); err != nil { http.Error(w, `{"error":"db"}`, 500); return }
    json.NewEncoder(w).Encode(map[string]string{"booking_no": bookingNo})
}

func (h *BookingHandler) GetByNo(w http.ResponseWriter, r *http.Request) {
    no := r.URL.Query().Get("booking_no")
    b, err := h.Repo.FindByNumber(r.Context(), no)
    if err != nil { http.Error(w, `{"error":"bulunamadı"}`, 404); return }
    json.NewEncoder(w).Encode(b)
}

func (h *BookingHandler) ListAdmin(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]interface{}{"data":[]})
}
func (h *BookingHandler) GetAdmin(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"detail":"booking"})
}
func (h *BookingHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status":"updated"})
}

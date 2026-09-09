package handler

import (
    "encoding/json"
    "net/http"
    "github.com/google/uuid"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/config"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/utils"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/repository"
)

type BookingCreateReq struct {
    TourID string `json:"tour_id"`
    DepartureID string `json:"departure_id"`
    AdultCount int `json:"adult_count"`
    ChildCount int `json:"child_count"`
    Travelers []struct {
        FullName string `json:"full_name"`
        BirthDate string `json:"birth_date"`
        IDNo string `json:"id_no"`
        IDType string `json:"id_type"`
    } `json:"travelers"`
}

func (h *BookingHandler) CreateFull(repo *repository.BookingRepo) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req BookingCreateReq
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, `{"error":"invalid"}`,400); return }
        cfg,_ := config.Load()
        // booking oluştur
        bookingID := uuid.NewString()
        bookingNo := "TA-"+uuid.NewString()[:8]
        // TODO: DB insert booking + travelers with encryption
        // encrypt IDNo
        for _,t := range req.Travelers {
            _, _ = utils.Encrypt(t.IDNo, cfg.EncryptionKey)
        }
        json.NewEncoder(w).Encode(map[string]string{"booking_no":bookingNo})
    }
}

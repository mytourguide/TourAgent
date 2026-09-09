package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/repository"
)

type TourHandler struct { Repo *repository.TourRepo }

func (h *TourHandler) List(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 1 { page = 1 }
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    if limit < 1 || limit > 100 { limit = 20 }

    tours, err := h.Repo.List(r.Context(), r.URL.Query().Get("category"), r.URL.Query().Get("location"), 0, 0, page, limit)
    if err != nil { http.Error(w, `{"error":"db error"}`, http.StatusInternalServerError); return }
    json.NewEncoder(w).Encode(map[string]interface{}{"data": tours})
}

func (h *TourHandler) GetBySlug(w http.ResponseWriter, r *http.Request) {
    slug := chi.URLParam(r, "slug")
    t, err := h.Repo.GetBySlug(r.Context(), slug)
    if err != nil { http.Error(w, `{"error":"bulunamadı"}`, http.StatusNotFound); return }
    json.NewEncoder(w).Encode(t)
}

func (h *TourHandler) Featured(w http.ResponseWriter, r *http.Request) {
    tours, err := h.Repo.Featured(r.Context(), 12)
    if err != nil { http.Error(w, `{"error":"db error"}`, http.StatusInternalServerError); return }
    json.NewEncoder(w).Encode(map[string]interface{}{"data": tours})
}

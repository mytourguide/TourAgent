package handler

import (
    "encoding/json"
    "net/http"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/config"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/auth"
)

type AuthRequest struct {
    Email string `json:"email"`
    Password string `json:"password"`
    FullName string `json:"full_name"`
}

func Register(w http.ResponseWriter, r *http.Request) {
    var req AuthRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, `{"error":"invalid"}`, 400); return }
    // Şifre hash ve kayıt DB'ye (basitleştirme)
    cfg, _ := config.Load()
    token, _ := auth.GenerateAccessToken("tmp-id", "customer", cfg.JWTSecret)
    json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

func Login(w http.ResponseWriter, r *http.Request) {
    var req AuthRequest
    json.NewDecoder(r.Body).Decode(&req)
    cfg, _ := config.Load()
    token, _ := auth.GenerateAccessToken("tmp-id", "customer", cfg.JWTSecret)
    json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

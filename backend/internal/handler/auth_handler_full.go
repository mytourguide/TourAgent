package handler

import (
    "encoding/json"
    "net/http"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/auth"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/config"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/repository"
)

type AuthHandler struct { UserRepo *repository.UserRepo }

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req struct {
        FullName string `json:"full_name"`
        Email string `json:"email"`
        Password string `json:"password"`
        Phone string `json:"phone"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil { http.Error(w, `{"error":"invalid"}`, 400); return }
    hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    id := uuid.NewString()
    if err := h.UserRepo.Create(r.Context(), id, req.FullName, req.Email, req.Phone, string(hash), "customer"); err != nil {
        http.Error(w, `{"error":"db"}`, 500); return
    }
    cfg,_ := config.Load()
    token,_ := auth.GenerateAccessToken(id, "customer", cfg.JWTSecret)
    json.NewEncoder(w).Encode(map[string]string{"access_token":token})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req struct { Email string `json:"email"`; Password string `json:"password"` }
    json.NewDecoder(r.Body).Decode(&req)
    user, err := h.UserRepo.FindByEmail(r.Context(), req.Email)
    if err != nil { http.Error(w, `{"error":"invalid credentials"}`, 401); return }
    if bcrypt.CompareHashAndPassword([]byte(user["password_hash"]), []byte(req.Password)) != nil {
        http.Error(w, `{"error":"invalid credentials"}`, 401); return
    }
    cfg,_ := config.Load()
    access,_ := auth.GenerateAccessToken(user["id"], user["role"], cfg.JWTSecret)
    refresh,_ := auth.GenerateRefreshToken(user["id"], cfg.JWTSecret)
    json.NewEncoder(w).Encode(map[string]string{"access_token":access,"refresh_token":refresh})
}

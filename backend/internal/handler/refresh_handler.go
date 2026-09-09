package handler

import (
    "encoding/json"
    "net/http"
    "github.com/golang-jwt/jwt/v5"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/auth"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/config"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
    var req struct{ RefreshToken string `json:"refresh_token"` }
    json.NewDecoder(r.Body).Decode(&req)
    cfg,_ := config.Load()
    claims := &jwt.RegisteredClaims{}
    _, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(t *jwt.Token)(interface{},error){return []byte(cfg.JWTSecret),nil})
    if err != nil { http.Error(w, `{"error":"invalid"}`,401); return }
    access,_ := auth.GenerateAccessToken(claims.Subject, "customer", cfg.JWTSecret)
    json.NewEncoder(w).Encode(map[string]string{"access_token":access})
}

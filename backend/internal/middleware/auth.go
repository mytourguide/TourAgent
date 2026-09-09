package middleware

import (
    "context"
    "net/http"
    "strings"

    "github.com/go-chi/chi/v5"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/auth"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/config"
)

type ctxKey string

const UserCtxKey ctxKey = "user"

func Auth(next http.Handler) http.Handler {
    cfg, _ := config.Load()
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        h := r.Header.Get("Authorization")
        if h == "" {
            http.Error(w, `{"error":"yetkisiz"}`, http.StatusUnauthorized)
            return
        }
        parts := strings.SplitN(h, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            http.Error(w, `{"error":"geçersiz token"}`, http.StatusUnauthorized)
            return
        }
        claims, err := auth.ParseToken(parts[1], cfg.JWTSecret)
        if err != nil {
            http.Error(w, `{"error":"geçersiz token"}`, http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), UserCtxKey, claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func AdminOnly(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        claims, ok := r.Context().Value(UserCtxKey).(*auth.Claims)
        if !ok || claims.Role != "admin" {
            http.Error(w, `{"error":"admin yetkisi gerekli"}`, http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func GetClaims(r *http.Request) *auth.Claims {
    if c, ok := r.Context().Value(UserCtxKey).(*auth.Claims); ok {
        return c
    }
    return nil
}

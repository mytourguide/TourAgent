package middleware

import (
    "net/http"
    "golang.org/x/time/rate"
    "sync"
)

var mu sync.Mutex
var clients = make(map[string]*rate.Limiter)

func getLimiter(ip string) *rate.Limiter {
    mu.Lock(); defer mu.Unlock()
    lim, ok := clients[ip]
    if !ok {
        lim = rate.NewLimiter(rate.Limit(2), 120)
        clients[ip] = lim
    }
    return lim
}

func RateLimit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        lim := getLimiter(r.RemoteAddr)
        if !lim.Allow() {
            http.Error(w, `{"error":"rate limit"}`, http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}

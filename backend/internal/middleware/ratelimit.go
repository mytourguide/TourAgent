package middleware

import (
    "net/http"
    "sync"
    "time"
)

type limiter struct {
    last    time.Time
    requests int
}

var mu sync.Mutex
var clients = make(map[string]*limiter)

func RateLimit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        mu.Lock()
        l, ok := clients[ip]
        if !ok {
            l = &limiter{}
            clients[ip] = l
        }
        mu.Unlock()
        now := time.Now()
        if now.Sub(l.last) > time.Minute {
            l.last = now
            l.requests = 0
        }
        l.requests++
        if l.requests > 120 {
            http.Error(w, `{"error":"rate limit"}`, http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}

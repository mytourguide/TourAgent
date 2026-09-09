package middleware

import (
    "encoding/json"
    "net/http"
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Basit örnek: istek gövdesini decode etmeye çalışmaz, sadece validator hazır.
        // Gerçek kullanımda handler'da struct'ı validate et.
        next.ServeHTTP(w, r)
    })
}

func ValidationError(w http.ResponseWriter, err error) {
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(map[string]interface{}{"error":err.Error()})
}

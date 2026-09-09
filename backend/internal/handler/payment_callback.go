package handler

import (
    "encoding/json"
    "net/http"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/email"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/config"
)

func IyzicoCallbackFull(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        Token string `json:"token"`
        PaymentStatus string `json:"paymentStatus"`
        ConversationID string `json:"conversationId"`
    }
    json.NewDecoder(r.Body).Decode(&payload)
    if payload.PaymentStatus == "SUCCESS" {
        // booking status -> paid
        // email gönder
        cfg,_ := config.Load()
        // email.Send(...)
        _ = email.Send("customer@example.com", "Ödeme Onayı", email.PaymentSuccess("TA-123"), email.SMTPConfig{Host:cfg.SMTPHost, Port:cfg.SMTPPort, User:cfg.SMTPUser, Pass:cfg.SMTPPass, From:cfg.SMTPFrom})
    }
    w.WriteHeader(http.StatusOK)
}

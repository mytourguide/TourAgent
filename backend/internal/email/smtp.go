package email

import (
    "fmt"
    "net/smtp"
    "strings"
)

type SMTPConfig struct {
    Host string
    Port int
    User string
    Pass string
    From string
}

func Send(to string, subject string, htmlBody string, cfg SMTPConfig) error {
    msg := []byte(fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=utf-8\r\n\r\n%s", to, cfg.From, subject, htmlBody))
    addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
    auth := smtp.PlainAuth("", cfg.User, cfg.Pass, cfg.Host)
    if err := smtp.SendMail(addr, auth, cfg.From, []string{to}, msg); err != nil {
        // Hata logla fakat request'i bozma
        return fmt.Errorf("mail gönderilemedi: %w", err)
    }
    return nil
}

func BookingConfirmation(to, bookingNo, tourTitle string) string {
    return fmt.Sprintf(`
    <html><body>
    <h2>Rezervasyonunuz Alındı</h2>
    <p>Merhaba,</p>
    <p>Rezervasyon numaranız <b>%s</b> ile tur <b>%s</b> için kaydınız oluşturuldu.</p>
    <p>Ödeme adımı tamamlandığında tekrar bilgi verilecektir.</p>
    </body></html>`, bookingNo, tourTitle)
}

func PaymentSuccess(to, bookingNo string) string {
    return fmt.Sprintf(`
    <html><body>
    <h2>Ödemeniz Başarıyla Alındı</h2>
    <p>Rezervasyon numaranız: <b>%s</b></p>
    <p>Turunuz onay süreci devam etmektedir.</p>
    </body></html>`, bookingNo)
}

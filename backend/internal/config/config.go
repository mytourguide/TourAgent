package config

import (
    "os"
    "strconv"
)

// App yapılandırma bilgilerini tutar. Tüm değerler environment variable'lardan yüklenir.
type Config struct {
    Port           string
    DatabaseURL    string
    JWTSecret      string
    EncryptionKey  string
    IyzicoAPIKey   string
    IyzicoSecret   string
    IyzicoBaseURL  string
    FrontendURL    string
    SMTPHost       string
    SMTPPort       int
    SMTPUser       string
    SMTPPass       string
    SMTPFrom       string
    S3Endpoint     string
    S3AccessKey    string
    S3SecretKey    string
    S3Bucket       string
    S3UsePathStyle bool
    RateLimit      int
}

func Load() (*Config, error) {
    // Basit yükleme; üretimde viper/override kullanılabilirdi.
    cfg := &Config{
        Port:           getEnv("PORT", "8080"),
        DatabaseURL:    getEnv("DATABASE_URL", "postgres://travel_user:change_me@localhost:5432/travel?sslmode=disable"),
        JWTSecret:      getEnv("JWT_SECRET", "change_me"),
        EncryptionKey:  getEnv("ENCRYPTION_KEY", ""),
        IyzicoAPIKey:   getEnv("IYZICO_API_KEY", ""),
        IyzicoSecret:   getEnv("IYZICO_SECRET_KEY", ""),
        IyzicoBaseURL:  getEnv("IYZICO_BASE_URL", "https://sandbox-api.iyzipay.com"),
        FrontendURL:    getEnv("FRONTEND_URL", "http://localhost:3000"),
        SMTPHost:       getEnv("SMTP_HOST", ""),
        SMTPUser:       getEnv("SMTP_USER", ""),
        SMTPPass:       getEnv("SMTP_PASS", ""),
        SMTPFrom:       getEnv("SMTP_FROM", "noreply@example.com"),
        S3Endpoint:     getEnv("S3_ENDPOINT", ""),
        S3AccessKey:    getEnv("S3_ACCESS_KEY", ""),
        S3SecretKey:    getEnv("S3_SECRET_KEY", ""),
        S3Bucket:       getEnv("S3_BUCKET", "tour-images"),
        S3UsePathStyle: getBoolEnv("S3_USE_PATH_STYLE", true),
        RateLimit:      getIntEnv("RATE_LIMIT_PER_MIN", 120),
    }
    if cfg.SMTPPort == 0 {
        cfg.SMTPPort = getIntEnv("SMTP_PORT", 587)
    }
    return cfg, nil
}

func getEnv(k, def string) string {
    if v := os.Getenv(k); v != "" {
        return v
    }
    return def
}
func getIntEnv(k string, def int) int {
    if v := os.Getenv(k); v != "" {
        if i, err := strconv.Atoi(v); err == nil {
            return i
        }
    }
    return def
}
func getBoolEnv(k string, def bool) bool {
    if v := os.Getenv(k); v != "" {
        return v == "true" || v == "1"
    }
    return def
}

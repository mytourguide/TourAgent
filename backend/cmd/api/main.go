package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/config"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/db"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/handler"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/middleware"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/repository"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("config yüklenemedi: %v", err)
    }

    ctx := context.Background()
    pool, err := db.NewPool(ctx, cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("db bağlantı hatası: %v", err)
    }
    defer pool.Close()

    if err := db.RunMigrations(ctx, pool); err != nil {
        log.Printf("migration uyarısı: %v", err)
    }

    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(60 * time.Second))
    // CORS basit
    r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", cfg.FrontendURL)
            w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type")
            if req.Method == http.MethodOptions {
                w.WriteHeader(http.StatusNoContent)
                return
            }
            next.ServeHTTP(w, req)
        })
    })

    r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })

    api := chi.NewRouter()
    api.Use(middleware.RateLimit)

    userRepo := repository.NewUserRepo(pool)
    authH := handler.AuthHandler{UserRepo: userRepo}
    // Public
    api.Post("/auth/register", authH.Register)
    api.Post("/auth/login", authH.Login)

    // Tours
    tourRepo := repository.NewTourRepo(pool)
    tourH := handler.TourHandler{Repo: tourRepo}
    api.Get("/tours", tourH.List)
    api.Get("/tours/featured", tourH.Featured)
    api.Get("/tours/{slug}", tourH.GetBySlug)

    // Bookings
    bookingRepo := repository.NewBookingRepo(pool)
    bookingH := handler.BookingHandler{Repo: bookingRepo}
    api.Post("/bookings", bookingH.Create)
    api.Get("/bookings", bookingH.GetByNo)

    // Payments
    api.Post("/payments/initiate", handler.InitPayment)
    api.Post("/payments/iyzico/callback", handler.IyzicoCallback)

    // Admin protected
    admin := chi.NewRouter()
    admin.Use(middleware.Auth)
    admin.Use(middleware.AdminOnly)

    itemsRepo := repository.NewTourItemsRepo(pool)
    admin.Get("/dashboard", handler.Dashboard)
    admin.Get("/tours", tourH.List)
    admin.Post("/tours", handler.AdminCreateTour)
    admin.Put("/tours/{id}", handler.AdminUpdateTour)
    admin.Delete("/tours/{id}", handler.AdminDeleteTour)
    admin.Put("/tours/{id}/items", handler.UpsertTourItems(itemsRepo))

    admin.Post("/tours/{id}/departures", handler.UpsertDepartures)
    admin.Put("/tours/{id}/departures/{departureId}", handler.UpdateDeparture)
    admin.Delete("/tours/{id}/departures/{departureId}", handler.DeleteDeparture)

    admin.Get("/bookings", bookingH.ListAdmin)
    admin.Get("/bookings/{id}", bookingH.GetAdmin)
    admin.Put("/bookings/{id}/status", bookingH.UpdateStatus)

    admin.Get("/payments", handler.ListPayments)
    admin.Post("/payments/{id}/refund", handler.RefundPayment)

    admin.Get("/categories", handler.ListCategories)
    admin.Post("/coupons", handler.CreateCoupon)
    admin.Get("/menus", handler.ListMenus)
    admin.Post("/menus", handler.CreateMenu)

    api.Mount("/admin", admin)

    r.Mount("/api", api)

    srv := &http.Server{
        Addr:    ":" + cfg.Port,
        Handler: r,
    }

    go func() {
        log.Printf("sunucu başlatılıyor :%s", cfg.Port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen hatası: %v", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("kapatılıyor...")
    ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    _ = srv.Shutdown(ctxShutdown)
    log.Println("sunucu durduruldu")
}

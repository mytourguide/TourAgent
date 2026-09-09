package models

import "time"

// Basit model tanımları — repository katmanı için. Tam alanlar ihtiyaca göre genişletilebilir.
type User struct {
    ID           string    `json:"id"`
    FullName     string    `json:"full_name"`
    Email        string    `json:"email"`
    Phone        string    `json:"phone"`
    Role         string    `json:"role"`
    EmailVerified bool     `json:"email_verified"`
    CreatedAt    time.Time `json:"created_at"`
}

type Tour struct {
    ID            string    `json:"id"`
    Title         string    `json:"title"`
    Slug          string    `json:"slug"`
    Description   string    `json:"description"`
    CategoryID    string    `json:"category_id"`
    CoverImage    string    `json:"cover_image"`
    DurationDays  int       `json:"duration_days"`
    DurationNights int     `json:"duration_nights"`
    Location      string    `json:"location"`
    BasePrice     float64   `json:"base_price"`
    Currency      string    `json:"currency"`
    IsActive      bool      `json:"is_active"`
    IsFeatured    bool      `json:"is_featured"`
}

type TourCategory struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Slug        string `json:"slug"`
    Description string `json:"description"`
}

type TourDeparture struct {
    ID             string    `json:"id"`
    TourID         string    `json:"tour_id"`
    StartDate      time.Time `json:"start_date"`
    EndDate        time.Time `json:"end_date"`
    Capacity       int       `json:"capacity"`
    FilledCapacity int       `json:"filled_capacity"`
    PricePerPerson *float64  `json:"price_per_person"`
    IsActive       bool      `json:"is_active"`
}

type Booking struct {
    ID            string    `json:"id"`
    BookingNo     string    `json:"booking_no"`
    TourID        string    `json:"tour_id"`
    DepartureID   string    `json:"departure_id"`
    UserID        string    `json:"user_id"`
    AdultCount    int       `json:"adult_count"`
    ChildCount    int       `json:"child_count"`
    TotalPrice    float64   `json:"total_price"`
    Currency      string    `json:"currency"`
    Status        string    `json:"status"`
    CreatedAt     time.Time `json:"created_at"`
}

type Payment struct {
    ID             string     `json:"id"`
    BookingID      string     `json:"booking_id"`
    IyzicoPaymentID string   `json:"iyzico_payment_id"`
    Amount         float64    `json:"amount"`
    Currency       string     `json:"currency"`
    Status         string     `json:"status"`
    Installments   int        `json:"installments"`
    PaidAt         *time.Time `json:"paid_at"`
    CreatedAt      time.Time  `json:"created_at"`
}

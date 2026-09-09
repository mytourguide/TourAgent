package repository

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/models"
)

type BookingRepo struct{ DB *pgxpool.Pool }

func NewBookingRepo(db *pgxpool.Pool) *BookingRepo { return &BookingRepo{DB: db} }

func (r *BookingRepo) Create(ctx context.Context, b models.Booking) error {
    query := `INSERT INTO bookings (id, booking_no, tour_id, departure_id, user_id, adult_count, child_count, total_price, currency, status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
    _, err := r.DB.Exec(ctx, query, b.ID, b.BookingNo, b.TourID, b.DepartureID, b.UserID, b.AdultCount, b.ChildCount, b.TotalPrice, b.Currency, b.Status)
    return err
}

func (r *BookingRepo) UpdateStatus(ctx context.Context, id string, status string) error {
    _, err := r.DB.Exec(ctx, `UPDATE bookings SET status=$1, updated_at=now() WHERE id=$2`, status, id)
    return err
}

func (r *BookingRepo) FindByNumber(ctx context.Context, no string) (*models.Booking, error) {
    var b models.Booking
    err := r.DB.QueryRow(ctx, `SELECT id, booking_no, tour_id, departure_id, user_id, adult_count, child_count, total_price, currency, status, created_at FROM bookings WHERE booking_no=$1`, no).Scan(&b.ID, &b.BookingNo, &b.TourID, &b.DepartureID, &b.UserID, &b.AdultCount, &b.ChildCount, &b.TotalPrice, &b.Currency, &b.Status, &b.CreatedAt)
    if err != nil { return nil, err }
    return &b, nil
}

package repository

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/models"
)

type TourRepo struct{ DB *pgxpool.Pool }

func NewTourRepo(db *pgxpool.Pool) *TourRepo { return &TourRepo{DB: db} }

func (r *TourRepo) List(ctx context.Context, categorySlug, loc string, minPrice, maxPrice float64, page, limit int) ([]models.Tour, error) {
    // Basit sorgu — filtreleme ilerletilebilir
    rows, err := r.DB.Query(ctx, `SELECT id, title, slug, description, category_id, cover_image, duration_days, duration_nights, location, base_price, currency, is_active, is_featured FROM tours WHERE is_active = true ORDER BY created_at DESC LIMIT $1 OFFSET $2`, limit, (page-1)*limit)
    if err != nil { return nil, err }
    defer rows.Close()
    var out []models.Tour
    for rows.Next() {
        var t models.Tour
        _ = rows.Scan(&t.ID, &t.Title, &t.Slug, &t.Description, &t.CategoryID, &t.CoverImage, &t.DurationDays, &t.DurationNights, &t.Location, &t.BasePrice, &t.Currency, &t.IsActive, &t.IsFeatured)
        out = append(out, t)
    }
    return out, nil
}

func (r *TourRepo) GetBySlug(ctx context.Context, slug string) (*models.Tour, error) {
    var t models.Tour
    err := r.DB.QueryRow(ctx, `SELECT id, title, slug, description, category_id, cover_image, duration_days, duration_nights, location, base_price, currency, is_active, is_featured FROM tours WHERE slug=$1`, slug).Scan(&t.ID, &t.Title, &t.Slug, &t.Description, &t.CategoryID, &t.CoverImage, &t.DurationDays, &t.DurationNights, &t.Location, &t.BasePrice, &t.Currency, &t.IsActive, &t.IsFeatured)
    if err != nil { return nil, err }
    return &t, nil
}

func (r *TourRepo) Featured(ctx context.Context, limit int) ([]models.Tour, error) {
    rows, err := r.DB.Query(ctx, `SELECT id, title, slug, description, category_id, cover_image, duration_days, duration_nights, location, base_price, currency, is_active, is_featured FROM tours WHERE is_active=true AND is_featured=true ORDER BY created_at DESC LIMIT $1`, limit)
    if err != nil { return nil, err }
    defer rows.Close()
    var out []models.Tour
    for rows.Next() { var t models.Tour; _=rows.Scan(&t.ID,&t.Title,&t.Slug,&t.Description,&t.CategoryID,&t.CoverImage,&t.DurationDays,&t.DurationNights,&t.Location,&t.BasePrice,&t.Currency,&t.IsActive,&t.IsFeatured); out=append(out,t) }
    return out, nil
}

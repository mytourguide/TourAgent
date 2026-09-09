package repository

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
)

type MenuRepo struct { DB *pgxpool.Pool }

func NewMenuRepo(db *pgxpool.Pool) *MenuRepo { return &MenuRepo{DB: db} }

func (r *MenuRepo) ListActive(ctx context.Context) ([]map[string]interface{}, error) {
    rows, err := r.DB.Query(ctx, `SELECT id, name, url, sort_order FROM site_menus WHERE is_active=true ORDER BY sort_order`)
    if err != nil { return nil, err }
    defer rows.Close()
    var out []map[string]interface{}
    for rows.Next() { /* scan */ }
    return out, nil
}

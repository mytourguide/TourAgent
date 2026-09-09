package repository

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct { DB *pgxpool.Pool }

func NewUserRepo(db *pgxpool.Pool) *UserRepo { return &UserRepo{DB: db} }

func (r *UserRepo) Create(ctx context.Context, id, fullName, email, phone, passwordHash, role string) error {
    _, err := r.DB.Exec(ctx, `INSERT INTO users(id, full_name, email, phone, password_hash, role) VALUES($1,$2,$3,$4,$5,$6)`, id, fullName, email, phone, passwordHash, role)
    return err
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (map[string]string, error) {
    row := r.DB.QueryRow(ctx, `SELECT id, password_hash, role, refresh_token FROM users WHERE email=$1`, email)
    var id, hash, role, refresh string
    if err := row.Scan(&id, &hash, &role, &refresh); err != nil { return nil, err }
    return map[string]string{"id":id,"password_hash":hash,"role":role,"refresh_token":refresh}, nil
}

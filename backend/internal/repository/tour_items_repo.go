package repository

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
)

type TourItemsRepo struct { DB *pgxpool.Pool }

func NewTourItemsRepo(db *pgxpool.Pool) *TourItemsRepo { return &TourItemsRepo{DB: db} }

// UpsertAll tek transaction içinde included/excluded/bring/itinerary/departures günceller
func (r *TourItemsRepo) UpsertAll(ctx context.Context, tourID string, included, excluded, bring []string, itinerary []struct{DayNo int; Title string; Description string}, departures []struct{StartDate string; EndDate string; Capacity int; Price float64}) error {
    tx, err := r.DB.BeginTx(ctx, nil)
    if err != nil { return err }
    defer tx.Rollback(ctx)

    // included items
    if _, err := tx.Exec(ctx, "DELETE FROM tour_included_items WHERE tour_id=$1", tourID); err != nil { return err }
    for i, v := range included {
        _, _ = tx.Exec(ctx, "INSERT INTO tour_included_items(tour_id, description, sort_order) VALUES($1,$2,$3)", tourID, v, i)
    }
    // excluded
    if _, err := tx.Exec(ctx, "DELETE FROM tour_excluded_items WHERE tour_id=$1", tourID); err != nil { return err }
    for i, v := range excluded {
        _, _ = tx.Exec(ctx, "INSERT INTO tour_excluded_items(tour_id, description, sort_order) VALUES($1,$2,$3)", tourID, v, i)
    }
    // bring
    if _, err := tx.Exec(ctx, "DELETE FROM tour_bring_items WHERE tour_id=$1", tourID); err != nil { return err }
    for i, v := range bring {
        _, _ = tx.Exec(ctx, "INSERT INTO tour_bring_items(tour_id, description, sort_order) VALUES($1,$2,$3)", tourID, v, i)
    }
    // itinerary
    if _, err := tx.Exec(ctx, "DELETE FROM tour_itinerary WHERE tour_id=$1", tourID); err != nil { return err }
    for _, it := range itinerary {
        _, _ = tx.Exec(ctx, "INSERT INTO tour_itinerary(tour_id, day_no, title, description) VALUES($1,$2,$3,$4)", tourID, it.DayNo, it.Title, it.Description)
    }
    // departures — mevcutları silmek yerine güncelleme stratejisi uygulanabilir
    // Basit yaklaşım: yeni gelenleri upsert
    for _, d := range departures {
        _, _ = tx.Exec(ctx, `INSERT INTO tour_departures(tour_id, start_date, end_date, capacity, price_per_person)
        VALUES($1,$2,$3,$4,$5)
        ON CONFLICT (tour_id, start_date) DO UPDATE SET end_date=$3, capacity=$4, price_per_person=$5`, tourID, d.StartDate, d.EndDate, d.Capacity, d.Price)
    }

    return tx.Commit(ctx)
}

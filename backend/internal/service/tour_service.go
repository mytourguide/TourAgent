package service

import (
    "context"
    "github.com/sizin-organizasyon/travel-agency/backend/internal/repository"
)

type TourService struct { Repo *repository.TourRepo }

func NewTourService(r *repository.TourRepo) *TourService { return &TourService{Repo: r} }

func (s *TourService) List(ctx context.Context, page, limit int) ([]interface{}, error) {
    tours, err := s.Repo.List(ctx, "", "", 0, 0, page, limit)
    if err != nil { return nil, err }
    return []interface{}(tours), nil
}

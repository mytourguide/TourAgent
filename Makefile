.PHONY: up down migrate seed tidy backend-run frontend-dev admin-dev test

up:
	docker compose up -d

down:
	docker compose down -v

migrate:
	psql "$${DATABASE_URL}" -f db/migrations/001_initial_schema.sql
	psql "$${DATABASE_URL}" -f db/migrations/002_seed.sql
	psql "$${DATABASE_URL}" -f db/migrations/003_admin_user.sql

seed:
	psql "$${DATABASE_URL}" -f db/migrations/002_seed.sql

tidy:
	cd backend && go mod tidy

backend-run:
	cd backend && go run ./cmd/api

frontend-dev:
	cd frontend && npm install && npm run dev -- -p 3000

admin-dev:
	cd admin && npm install && npm run dev -- -p 3001

test:
	cd backend && go test ./...

# Seyahat Acentası — Kurulum Kılavuzu

## İyileştirmeler
- JWT + bcrypt + refresh token
- Validator middleware
- Service katmanı eklendi
- iyzico imza üretimi
- Payment callback + e-posta
- Frontend API fetch
- Admin login sayfası
- Docker compose tam servis seti

# Seyahat Acentası — Kurulum Kılavuzu

Bu proje Monorepo yapısındadır:
- backend/  → Go + chi + PostgreSQL + iyzico
- frontend/ → Next.js müşteri sitesi
- admin/    → Next.js yönetim paneli

## Ön Koşullar
- Docker ve docker-compose
- Node 18+
- Go 1.22+ (backend geliştirmek için)
- PostgreSQL istemcisi (psql) opsiyonel

## Hızlı Başlangıç
```bash
# depolara taşı
cd travel-agency

# altyapıyı başlat
docker compose up -d postgres minio

# DB şemalarını uygula
export DATABASE_URL=postgres://travel_user:change_me@localhost:5432/travel?sslmode=disable
psql "$DATABASE_URL" -f db/migrations/001_initial_schema.sql
psql "$DATABASE_URL" -f db/migrations/002_seed.sql
psql "$DATABASE_URL" -f db/migrations/003_admin_user.sql

# Backend
cd backend
cp ../.env.example .env
go mod tidy
go run ./cmd/api

# Frontend
cd ../frontend
npm install
npm run dev -- -p 3000

# Admin
cd ../admin
npm install
npm run dev -- -p 3001
```

## iyzico Sandbox
Sandbox anahtarlarını .env'e koy. Test kartları iyzico dokümantasyonunda.
Kart bilgileri asla sunucuya ulaşmaz; checkout form embed edilir.

## Yönetici Kullanıcısı
003_admin_user.sql içindeki password_hash placebo'dur. Gerçek admin oluşturmak için:
```go
bytes := []byte("admin123")
hash, _ := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
```
elde edilen hash'i SQL'e koy.

## Mimari
Katmanlı Go backend (handler→service→repository). JWT access+refresh. 
Frontend SSR/SSG ile SEO. Admin Next.js route group.
KVKK için id_no şifrelemesi (AES-GCM). S3 presigned URL.

-- ============================================================
-- 001_initial_schema.sql
-- Seyahat Acentası veritabanı şeması — PostgreSQL
-- Not: UUID'ler için gen yerleşik `gen_random_uuid()` kullanılır
-- (PostgreSQL 13+). pgcrypto gerekmez.
-- ============================================================

BEGIN;

-- ------------------------------------------------------------
-- UZANTILAR
-- ------------------------------------------------------------
-- pgcrypto sadece eski PG için; modern PG'de gen_random_uuid yerleşik.
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- ------------------------------------------------------------
-- KULLANICILAR (users)
-- ------------------------------------------------------------
-- Rol: 'customer' (müşteri) | 'admin' (yönetici)
CREATE TABLE users (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name         VARCHAR(150) NOT NULL,
    email             VARCHAR(255) NOT NULL UNIQUE,
    phone             VARCHAR(32),
    password_hash     VARCHAR(255) NOT NULL,
    role              VARCHAR(20)  NOT NULL DEFAULT 'customer'
                      CHECK (role IN ('customer', 'admin')),
    email_verified    BOOLEAN NOT NULL DEFAULT FALSE,
    refresh_token     TEXT,            -- şifrelenmiş refresh token
    two_fa_secret     TEXT,            -- opsiyonel TOTP sırrı (admin 2FA)
    two_fa_enabled    BOOLEAN NOT NULL DEFAULT FALSE,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_role  ON users (role);

-- ------------------------------------------------------------
-- TUR KATEGORİLERİ (tour_categories)
-- ------------------------------------------------------------
CREATE TABLE tour_categories (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(100) NOT NULL,
    slug        VARCHAR(120) NOT NULL UNIQUE,
    description TEXT,
    sort_order  INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ------------------------------------------------------------
-- TURLAR (tours)
-- ------------------------------------------------------------
CREATE TABLE tours (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title           VARCHAR(200) NOT NULL,
    slug            VARCHAR(220) NOT NULL UNIQUE,
    description     TEXT,
    category_id     UUID REFERENCES tour_categories(id) ON DELETE SET NULL,
    cover_image     TEXT,             -- S3 object key veya tam URL
    duration_days   INTEGER NOT NULL DEFAULT 1, -- gün
    duration_nights INTEGER NOT NULL DEFAULT 0, -- gece
    location        VARCHAR(200),
    base_price      NUMERIC(12,2) NOT NULL DEFAULT 0, -- taban fiyat (kur parası)
    currency        VARCHAR(3)  NOT NULL DEFAULT 'TRY',
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    is_featured     BOOLEAN NOT NULL DEFAULT FALSE, -- ana sayfa öne çıkan
    meta_title      VARCHAR(200),      -- SEO
    meta_description TEXT,             -- SEO
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_tours_category ON tours (category_id);
CREATE INDEX idx_tours_active   ON tours (is_active);
CREATE INDEX idx_tours_slug     ON tours (slug);

-- ------------------------------------------------------------
-- TUR GÖRSELLERİ (tour_images)
-- ------------------------------------------------------------
CREATE TABLE tour_images (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tour_id    UUID NOT NULL REFERENCES tours(id) ON DELETE CASCADE,
    image_url  TEXT NOT NULL,
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_tour_images_tour ON tour_images (tour_id);

-- ------------------------------------------------------------
-- KALKIŞ TARİHLERİ (tour_departures)
-- ------------------------------------------------------------
-- Tarih bazlı fiyat/kontenjan farkı için ayrı tablo.
CREATE TABLE tour_departures (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tour_id         UUID NOT NULL REFERENCES tours(id) ON DELETE CASCADE,
    start_date      DATE NOT NULL,
    end_date        DATE NOT NULL,
    capacity        INTEGER NOT NULL DEFAULT 0,       -- toplam kontenjan
    filled_capacity INTEGER NOT NULL DEFAULT 0,       -- dolu kontenjan
    price_per_person NUMERIC(12,2),                   -- NULL ise tour.base_price kullanılır
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_departures_tour  ON tour_departures (tour_id);
CREATE INDEX idx_departures_date  ON tour_departures (start_date);
CREATE UNIQUE INDEX uq_departure_tour_date ON tour_departures (tour_id, start_date);

-- ------------------------------------------------------------
-- TURA DAHİL OLANLAR (tour_included_items)
-- ------------------------------------------------------------
CREATE TABLE tour_included_items (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tour_id     UUID NOT NULL REFERENCES tours(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    sort_order  INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_included_tour ON tour_included_items (tour_id);

-- ------------------------------------------------------------
-- TURA DAHİL OLMAYANLAR (tour_excluded_items)
-- ------------------------------------------------------------
CREATE TABLE tour_excluded_items (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tour_id     UUID NOT NULL REFERENCES tours(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    sort_order  INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_excluded_tour ON tour_excluded_items (tour_id);

-- ------------------------------------------------------------
-- YANINIZDA GETİRMENİZ GEREKENLER (tour_bring_items)
-- ------------------------------------------------------------
CREATE TABLE tour_bring_items (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tour_id     UUID NOT NULL REFERENCES tours(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    sort_order  INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_bring_tour ON tour_bring_items (tour_id);

-- ------------------------------------------------------------
-- GÜN GÜN PROGRAM (tour_itinerary)
-- ------------------------------------------------------------
CREATE TABLE tour_itinerary (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tour_id     UUID NOT NULL REFERENCES tours(id) ON DELETE CASCADE,
    day_no      INTEGER NOT NULL,
    title       VARCHAR(200),
    description TEXT,
    sort_order  INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_itinerary_tour ON tour_itinerary (tour_id);

-- ------------------------------------------------------------
-- KUPONLAR (coupons)
-- ------------------------------------------------------------
-- indirim_tipi: 'percentage' (yüzde) | 'fixed' (sabit tutar)
CREATE TABLE coupons (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code          VARCHAR(50) NOT NULL UNIQUE,
    discount_type VARCHAR(20) NOT NULL DEFAULT 'percentage'
                  CHECK (discount_type IN ('percentage', 'fixed')),
    discount_value NUMERIC(12,2) NOT NULL,
    valid_from    DATE,
    valid_until   DATE,
    max_uses      INTEGER,
    used_count    INTEGER NOT NULL DEFAULT 0,
    is_active     BOOLEAN NOT NULL DEFAULT TRUE,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ------------------------------------------------------------
-- REZERVASYONLAR (bookings)
-- ------------------------------------------------------------
-- Durumlar: pending | paid | confirmed | cancelled | failed
CREATE TABLE bookings (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_no        VARCHAR(20) NOT NULL UNIQUE,   -- müşteriye gösterilen numara
    tour_id           UUID NOT NULL REFERENCES tours(id),
    departure_id      UUID NOT NULL REFERENCES tour_departures(id),
    user_id           UUID REFERENCES users(id) ON DELETE SET NULL,
    adult_count       INTEGER NOT NULL DEFAULT 0,
    child_count       INTEGER NOT NULL DEFAULT 0,
    total_price       NUMERIC(12,2) NOT NULL,
    currency          VARCHAR(3) NOT NULL DEFAULT 'TRY',
    status            VARCHAR(20) NOT NULL DEFAULT 'pending'
                      CHECK (status IN ('pending', 'paid', 'confirmed',
                                        'cancelled', 'failed')),
    coupon_id         UUID REFERENCES coupons(id),   -- kupon tablosu aşağıda
    discount_amount   NUMERIC(12,2) NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_bookings_tour        ON bookings (tour_id);
CREATE INDEX idx_bookings_departure   ON bookings (departure_id);
CREATE INDEX idx_bookings_user        ON bookings (user_id);
CREATE INDEX idx_bookings_status      ON bookings (status);
CREATE INDEX idx_bookings_created     ON bookings (created_at);

-- ------------------------------------------------------------
-- YOLCU BİLGİLERİ (booking_travelers)
-- ------------------------------------------------------------
-- KVKK: ad/soyad ve doğum tarihi düz; kimlik/pasaport no application
-- seviyesinde AES-GCM ile şifrelenerek saklanır (id_no_encrypted).
CREATE TABLE booking_travelers (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id       UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    full_name        VARCHAR(150) NOT NULL,
    birth_date       DATE,
    id_no_encrypted  TEXT,            -- şifreli kimlik/pasaport no
    id_type          VARCHAR(20),     -- 'national_id' | 'passport'
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_travelers_booking ON booking_travelers (booking_id);

-- ------------------------------------------------------------
-- ÖDEMELER (payments)
-- ------------------------------------------------------------
-- Kart verisi ASLA saklanmaz; sadece iyzico referansları tutulur.
-- Durum: pending | paid | failed | refunded | partially_refunded
CREATE TABLE payments (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id        UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    iyzico_payment_id VARCHAR(100),
    iyzico_conversation_id VARCHAR(100),
    amount            NUMERIC(12,2) NOT NULL,
    currency          VARCHAR(3) NOT NULL DEFAULT 'TRY',
    status            VARCHAR(30) NOT NULL DEFAULT 'pending'
                      CHECK (status IN ('pending', 'paid', 'failed',
                                        'refunded', 'partially_refunded')),
    installments      INTEGER NOT NULL DEFAULT 1, -- 1 = peşin, 2/3/6/9 taksit
    paid_at           TIMESTAMPTZ,
    refunded_at       TIMESTAMPTZ,
    raw_response      JSONB,          -- iyzico callback ham verisi (log)
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_payments_booking ON payments (booking_id);
CREATE INDEX idx_payments_status  ON payments (status);

-- ------------------------------------------------------------
-- İADELER (refunds)
-- ------------------------------------------------------------
CREATE TABLE refunds (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    payment_id        UUID NOT NULL REFERENCES payments(id) ON DELETE CASCADE,
    amount            NUMERIC(12,2) NOT NULL,
    reason            TEXT,
    status            VARCHAR(20) NOT NULL DEFAULT 'pending'
                      CHECK (status IN ('pending', 'processed', 'failed')),
    iyzico_refund_id  VARCHAR(100),
    initiated_by      UUID REFERENCES users(id) ON DELETE SET NULL, -- admin
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_refunds_payment ON refunds (payment_id);

-- ------------------------------------------------------------
-- YOLCULUK DURUM LOGLARI (booking_status_logs)
-- ------------------------------------------------------------
CREATE TABLE booking_status_logs (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    from_status VARCHAR(20),
    to_status  VARCHAR(20) NOT NULL,
    note       TEXT,
    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_statuslog_booking ON booking_status_logs (booking_id);

-- ------------------------------------------------------------
-- trigger: users.updated_at / tours.updated_at / bookings.updated_at
-- otomatik güncelleme fonksiyonu
-- ------------------------------------------------------------
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_updated      BEFORE UPDATE ON users      FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_tours_updated      BEFORE UPDATE ON tours      FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_bookings_updated   BEFORE UPDATE ON bookings   FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_payments_updated   BEFORE UPDATE ON payments   FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_refunds_updated    BEFORE UPDATE ON refunds    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

COMMIT;
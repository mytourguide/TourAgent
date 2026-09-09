-- ============================================================
-- 002_seed.sql — Örnek veri (development)
-- ÜRETİMDE ÇALIŞTIRMAYIN.
-- ============================================================

BEGIN;

-- Kategoriler
INSERT INTO tour_categories (id, name, slug, description, sort_order) VALUES
 (gen_random_uuid(), 'Yurt İçi Turlar',   'yurt-ici',    'Türkiye içi turlar', 1),
 (gen_random_uuid(), 'Yurt Dışı Turlar',  'yurt-disi',   'Uluslararası destinasyonlar', 2),
 (gen_random_uuid(), 'Kültür Turları',    'kultur',      'Tarih ve kültür odaklı', 3),
 (gen_random_uuid(), 'Doğa Turları',      'doga',        'Doğa aktiviteleri', 4),
 (gen_random_uuid(), 'Balayı Paketleri',  'balayi',      'Romantik kaçışlar', 5),
 (gen_random_uuid(), 'Mavi Turlar',       'mavi-turlar', 'Ege ve Akdeniz kıyıları tekne turları', 6),
 (gen_random_uuid(), 'Extrem Turlar',     'extrem-turlar','Aşırı macera ve adrenalin', 7),
 (gen_random_uuid(), 'Aile Turları',      'aile-turlar', 'Çocuk dostu aile paketleri', 8),
 (gen_random_uuid(), 'Macera Turları',    'macera-turlar','Yürüyüş, rafting, kamp', 9);

-- Demo Sadece şema testi — gerçek tur ekleme admin panelinden yapılmalı.
COMMIT;
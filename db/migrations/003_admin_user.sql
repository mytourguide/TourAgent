-- ------------------------------------------------------------
-- 003_admin_user.sql
-- Örnek admin kullanıcısı — ÜRETİMDE düzelt!
-- password_hash: bcrypt ile oluşturulmuş hash olmalı.
-- Örnek üretim:
--   hash, _ := bcrypt.GenerateFromPassword([]byte("güçlüŞifre123!"), bcrypt.DefaultCost)
-- ------------------------------------------------------------
INSERT INTO users (id, full_name, email, phone, password_hash, role, email_verified)
VALUES (
  gen_random_uuid(),
  'Admin Kullanıcı',
  'admin@example.com',
  '+90 5xx xxx xx xx',
  '$2a$10$PLACEHOLDERHASHREPLACEWITHREALBCRYPT', -- DEĞİŞTİR
  'admin',
  TRUE
)
ON CONFLICT (email) DO NOTHING;

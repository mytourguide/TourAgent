INSERT INTO users (id, full_name, email, phone, password_hash, role, email_verified)
VALUES (
  gen_random_uuid(),
  'Admin Kullanıcı',
  'admin@example.com',
  '+90 5xx xxx xx xx',
  '$2b$12$9/d3wleDvVW2Idio/1CK6.orPwuYteAWCT2jpmu65HjEJmY/UYFcm',
  'admin',
  TRUE
)
ON CONFLICT (email) DO NOTHING;

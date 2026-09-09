# Veritabanı Migrations

Sıra:
1. 001_initial_schema.sql — tüm tablolar
2. 002_seed.sql — örnek kategoriler
3. 003_admin_user.sql — admin kullanıcı

Uygulama:
```bash
psql "$DATABASE_URL" -f db/migrations/001_initial_schema.sql
...
```
Backend başlatılırken otomatik migration çalıştırılabilir.
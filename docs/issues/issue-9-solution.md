# Issue #9 — Add Postgres support

## What was done

Added a Postgres repository implementation and made the agenda service switchable via `DB_DRIVER=postgres` without changing business logic.

### New package: `services/agenda/internal/repository/postgres`

- **db.go:** `Open(dsn string)` using `gorm.io/driver/postgres`, `Migrate(db)` with the same models as SQLite (Doctor, WorkingHours, Patient, Reservation, BlockedSlot).
- **doctor.go, patient.go, reservation.go, blocked_slot.go:** Same method signatures and behavior as SQLite; implement domain repository interfaces.
- **seed.go:** Same seed logic as SQLite (doctors, working hours, blocked slots) so a fresh Postgres DB gets default data.

### Agenda main

- **openDB:** When `cfg.DBDriver == "postgres"`, call `postgres.Open(cfg.DBSource)`; when `"sqlite3"`, keep existing `sqlite.Open(cfg.DBSource)`.
- After opening: switch on driver to call the matching `Migrate` and `Seed`.
- Construct repositories with the same domain interfaces; switch on driver to use `sqlite.New*` or `postgres.New*`. Usecases and gRPC server are unchanged.

### Config

- Existing `DB_DRIVER` (default `sqlite3`) and `DB_SOURCE` (or `DATABASE_URL`) work; for Postgres, set `DB_DRIVER=postgres` and `DB_SOURCE` to a DSN.

### Dependency

- **go.mod:** Added `gorm.io/driver/postgres v1.6.0`.

Business logic stays in domain/usecase; only a second repository implementation was added.

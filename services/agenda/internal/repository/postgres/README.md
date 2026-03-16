# Postgres repository

Implements the same domain repository interfaces as the SQLite package. Use with `DB_DRIVER=postgres` and `DB_SOURCE` set to a Postgres DSN (e.g. `host=localhost user=agenda password=secret dbname=agenda sslmode=disable`).

Migrate and seed run automatically on startup when using the postgres driver.

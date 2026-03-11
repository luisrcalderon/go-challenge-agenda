package sqlite_test

import (
	"context"
	"runtime"
	"testing"

	"go-challenge-agenda/services/agenda/internal/repository/sqlite"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func openTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := sqlite.Open(":memory:")
	require.NoError(t, err)
	require.NoError(t, sqlite.Migrate(db))
	require.NoError(t, sqlite.Seed(db))
	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})
	return db
}

// TestListDoctors_NoGoroutineLeak verifies that repeated calls to ListDoctors
// do not leak goroutines (e.g. from unclosed rows).
func TestListDoctors_NoGoroutineLeak(t *testing.T) {
	db := openTestDB(t)
	repo := sqlite.NewDoctorRepository(db)

	// Warm up
	_, _ = repo.ListDoctors(context.Background())

	before := runtime.NumGoroutine()

	const iterations = 100
	for range iterations {
		_, err := repo.ListDoctors(context.Background())
		require.NoError(t, err)
	}

	runtime.GC()
	after := runtime.NumGoroutine()

	// Allow a small tolerance for runtime goroutines
	leaked := after - before
	assert.LessOrEqual(t, leaked, 2,
		"ListDoctors leaked goroutines: before=%d after=%d delta=%d", before, after, leaked)
}

// TestListDoctors_WorkingHoursNotEmpty verifies that GetDoctor returns working hours
// while ListDoctors exposes the missing Preload bug (working_hours will be nil/empty).
// This test documents the behavioral difference between the two methods.
func TestListDoctors_WorkingHoursNotEmpty(t *testing.T) {
	db := openTestDB(t)
	repo := sqlite.NewDoctorRepository(db)

	single, err := repo.GetDoctor(context.Background(), "doc-001")
	require.NoError(t, err)
	assert.NotEmpty(t, single.WorkingHours, "GetDoctor should return working hours")

	list, err := repo.ListDoctors(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, list)

	// This assertion FAILS due to the missing Preload("WorkingHours") bug in ListDoctors.
	assert.NotEmpty(t, list[0].WorkingHours,
		"ListDoctors should return working hours for each doctor")
}

// BenchmarkListDoctors measures allocation and throughput of the list operation.
func BenchmarkListDoctors(b *testing.B) {
	db, err := sqlite.Open(":memory:")
	require.NoError(b, err)
	require.NoError(b, sqlite.Migrate(db))
	require.NoError(b, sqlite.Seed(db))
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo := sqlite.NewDoctorRepository(db)
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		_, _ = repo.ListDoctors(ctx)
	}
}

// BenchmarkGetDoctor measures single-record fetch with Preload.
func BenchmarkGetDoctor(b *testing.B) {
	db, err := sqlite.Open(":memory:")
	require.NoError(b, err)
	require.NoError(b, sqlite.Migrate(db))
	require.NoError(b, sqlite.Seed(db))
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo := sqlite.NewDoctorRepository(db)
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		_, _ = repo.GetDoctor(ctx, "doc-001")
	}
}

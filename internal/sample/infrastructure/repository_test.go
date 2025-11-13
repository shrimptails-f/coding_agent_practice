package infrastructure

import (
	"testing"
	"time"

	"business/internal/library/mysql"
	"business/internal/sample/domain"
	"business/tools/migrations/model"
)

func prepareTestRepo(t *testing.T) *SampleRepository {
	t.Helper()

	conn, err := mysql.NewTest()
	if err != nil {
		t.Fatalf("failed to init test mysql: %v", err)
	}

	migrationTargets := []interface{}{model.Sample{}}
	if err := conn.DB.Migrator().DropTable(migrationTargets...); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}

	if err := conn.DB.AutoMigrate(&SampleRecord{}); err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}

	t.Cleanup(func() {
		conn.DB.Migrator().DropTable(migrationTargets...)
	})

	return NewRepository(conn.DB)
}

func TestSampleRepository_ListOrdersDescending(t *testing.T) {
	repo := prepareTestRepo(t)

	older := domain.Sample{
		Title:     "older",
		CreatedAt: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	newer := domain.Sample{
		Title:     "newer",
		CreatedAt: time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
	}

	for _, s := range []domain.Sample{older, newer} {
		if err := repo.Save(s); err != nil {
			t.Fatalf("save failed: %v", err)
		}
	}

	got, err := repo.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("expected 2 samples, got %d", len(got))
	}

	if got[0].Title != newer.Title || got[1].Title != older.Title {
		t.Fatalf("expected desc order, got %+v", got)
	}
}

func TestSampleRepository_SaveIgnoresDuplicateTitles(t *testing.T) {
	repo := prepareTestRepo(t)

	first := domain.Sample{
		Title:     "dup",
		CreatedAt: time.Date(2024, time.February, 1, 0, 0, 0, 0, time.UTC),
	}

	second := domain.Sample{
		Title:     "dup",
		CreatedAt: time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
	}

	if err := repo.Save(first); err != nil {
		t.Fatalf("first save failed: %v", err)
	}

	if err := repo.Save(second); err != nil {
		t.Fatalf("second save failed: %v", err)
	}

	got, err := repo.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}

	if len(got) != 1 {
		t.Fatalf("expected 1 sample stored, got %d", len(got))
	}

	if !got[0].CreatedAt.Equal(first.CreatedAt) {
		t.Fatalf("expected first record unchanged, got %+v", got[0])
	}
}

package application

import (
	"business/internal/sample/domain"
	"errors"
	"testing"
)

type stubRepo struct {
	saved []domain.Sample
	list  []domain.Sample
	err   error
}

func (s *stubRepo) List() ([]domain.Sample, error) {
	return s.list, s.err
}

func (s *stubRepo) Save(sample domain.Sample) error {
	if s.err != nil {
		return s.err
	}
	s.saved = append(s.saved, sample)
	return nil
}

func TestCreateSample(t *testing.T) {
	repo := &stubRepo{}
	uc := NewUseCase(repo)

	title := "Hello"

	sample, err := uc.CreateSample(domain.CreateSampleInput{Title: title})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sample.Title != title {
		t.Fatalf("expected %s", title)
	}

	if len(repo.saved) != 1 {
		t.Fatalf("expected repository to save 1 sample, got %d", len(repo.saved))
	}
}

func TestCreateSample_ValidationError(t *testing.T) {
	repo := &stubRepo{}
	uc := NewUseCase(repo)

	if _, err := uc.CreateSample(domain.CreateSampleInput{}); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestCreateSample_RepositoryError(t *testing.T) {
	repo := &stubRepo{err: errors.New("boom")}
	uc := NewUseCase(repo)

	if _, err := uc.CreateSample(domain.CreateSampleInput{Title: "ok"}); err == nil {
		t.Fatal("expected repository error")
	}
}

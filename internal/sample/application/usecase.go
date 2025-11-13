package application

import (
	"business/internal/sample/domain"
)

// Repository expresses the minimal persistence contract required by the use case.
type Repository interface {
	List() ([]domain.Sample, error)
	Save(sample domain.Sample) error
}

// UseCase exposes the sample-oriented operations used by the presentation layer.
type UseCase interface {
	ListSamples() ([]domain.Sample, error)
	CreateSample(input domain.CreateSampleInput) (domain.Sample, error)
}

type sampleUseCase struct {
	repo Repository
}

// NewUseCase wires the dependencies together.
func NewUseCase(repo Repository) UseCase {
	return &sampleUseCase{repo: repo}
}

func (uc *sampleUseCase) ListSamples() ([]domain.Sample, error) {
	return uc.repo.List()
}

func (uc *sampleUseCase) CreateSample(input domain.CreateSampleInput) (domain.Sample, error) {
	if err := input.Validate(); err != nil {
		return domain.Sample{}, err
	}

	sample := domain.NewSample(input)
	if err := uc.repo.Save(sample); err != nil {
		return domain.Sample{}, err
	}
	return sample, nil
}

package infrastructure

import (
	"business/internal/sample/application"
	"business/internal/sample/domain"
	"time"

	"gorm.io/gorm"
)

var _ application.Repository = (*SampleRepository)(nil)

// SampleRepository persists samples into the samplese table via GORM.
type SampleRepository struct {
	db *gorm.DB
}

// NewRepository wraps the provided gorm DB. All tables are assumed to exist beforehand.
func NewRepository(db *gorm.DB) *SampleRepository {
	if db == nil {
		panic("gorm db is nil")
	}
	return &SampleRepository{db: db}
}

func (r *SampleRepository) List() ([]domain.Sample, error) {
	var records []SampleRecord
	if err := r.db.Order("created_at desc").Find(&records).Error; err != nil {
		return nil, err
	}

	samples := make([]domain.Sample, len(records))
	for i, rec := range records {
		samples[i] = rec.toDomain()
	}
	return samples, nil
}

func (r *SampleRepository) Save(sample domain.Sample) error {
	record := newSampleRecord(sample)
	return r.db.
		Where(&SampleRecord{Title: record.Title}).
		FirstOrCreate(&record).Error
}

// SampleRecord is the GORM model mapped to the "samplese" table.
type SampleRecord struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (SampleRecord) TableName() string {
	return "samplese"
}

func newSampleRecord(sample domain.Sample) SampleRecord {
	return SampleRecord{
		Title:     sample.Title,
		CreatedAt: sample.CreatedAt,
	}
}

func (s SampleRecord) toDomain() domain.Sample {
	return domain.Sample{
		ID:        s.ID,
		Title:     s.Title,
		CreatedAt: s.CreatedAt,
	}
}

package seeders

import (
	"business/tools/migrations/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateSamples inserts sample rows into samplese table.
func CreateSamples(tx *gorm.DB) error {
	now := time.Now()
	samples := []model.Sample{
		{Title: "First Sample", CreatedAt: now},
		{Title: "Second Sample", CreatedAt: now},
	}

	for _, s := range samples {
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&s).Error; err != nil {
			return err
		}
	}
	return nil
}

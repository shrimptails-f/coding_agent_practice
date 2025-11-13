package model

import "time"

// Sample represents the table structure for samplese.
type Sample struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Sample) TableName() string {
	return "samplese"
}

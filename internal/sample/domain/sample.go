package domain

import (
	"errors"
	"strings"
	"time"
)

// Sample represents the smallest unit of data flowing through the sample stack.
type Sample struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateSampleInput carries the payload coming from the presentation layer.
type CreateSampleInput struct {
	Title string
}

// Validate performs minimal sanity checks so the application layer can trust the data.
func (in CreateSampleInput) Validate() error {
	title := strings.TrimSpace(in.Title)
	if title == "" {
		return errors.New("title is required")
	}
	if len([]rune(title)) > 100 {
		return errors.New("title must be 100 characters or less")
	}
	return nil
}

// NewSample is a tiny factory to keep construction logic in one place.
func NewSample(input CreateSampleInput) Sample {
	return Sample{
		Title:     strings.TrimSpace(input.Title),
		CreatedAt: time.Now(),
	}
}

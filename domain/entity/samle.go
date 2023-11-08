package entity

import "time"

type Sample struct {
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func (s *Sample) Id() string {
	return s.id
}

func (s *Sample) Name() string {
	return s.name
}

func (s *Sample) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Sample) UpdatedAt() time.Time {
	return s.updatedAt
}

func NewSample(id string, name string, createdAt time.Time, updatedAt time.Time) Sample {
	return Sample{
		id:        id,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

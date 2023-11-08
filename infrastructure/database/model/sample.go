package model

import (
	"time"

	"kiravia.com/internship-go-api/domain/entity"
)

type Sample struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Sample) ToEntity() entity.Sample {
	return entity.NewSample(s.Id, s.Name, s.CreatedAt, s.UpdatedAt)
}

func NewSampleFromEntity(sample entity.Sample) Sample {
	return Sample{
		Id:        sample.Id(),
		Name:      sample.Name(),
		CreatedAt: sample.CreatedAt(),
		UpdatedAt: sample.UpdatedAt(),
	}
}

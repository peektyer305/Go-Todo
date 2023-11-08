package database

import (
	"context"

	"github.com/samber/lo"
	"gorm.io/gorm"
	"kiravia.com/internship-go-api/domain"
	"kiravia.com/internship-go-api/domain/entity"
	"kiravia.com/internship-go-api/infrastructure/database/model"
)

type SampleRepository struct {
	db *gorm.DB
}

func (s SampleRepository) Save(ctx context.Context, sample entity.Sample) (*entity.Sample, error) {
	conn := s.db.WithContext(ctx)
	sampleModel := model.NewSampleFromEntity(sample)
	if err := conn.Save(&sampleModel).Error; err != nil {
		return nil, err
	}
	return lo.ToPtr(sampleModel.ToEntity()), nil
}

func (s SampleRepository) FindById(ctx context.Context, id string) (*entity.Sample, error) {
	conn := s.db.WithContext(ctx)
	var sampleModel model.Sample
	if err := conn.Where("id = ?", id).First(&sampleModel).Error; err != nil {
		return nil, err
	}
	return lo.ToPtr(sampleModel.ToEntity()), nil
}

func NewSampleRepository(db *gorm.DB) domain.ISampleRepository {
	return &SampleRepository{db: db}
}

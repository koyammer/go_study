package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IGachaRepository interface {
	GetGachaById(gacha *model.Gacha, id string) error
	CreateGacha(gacha *model.Gacha) error
}

type gachaRepository struct {
	db *gorm.DB
}

func NewGachaRepository(db *gorm.DB) IGachaRepository {
	return &gachaRepository{db}
}

func (gr *gachaRepository) GetGachaById(gacha *model.Gacha, id string) error {
	if err := gr.db.Where("id=?", id).First(gacha).Error; err != nil {
		return err
	}
	return nil
}

func (gr *gachaRepository) CreateGacha(gacha *model.Gacha) error {
	if err := gr.db.Create(gacha).Error; err != nil {
		return err
	}
	return nil
}

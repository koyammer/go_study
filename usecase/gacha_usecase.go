package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IGachaUsecase interface {
	GetGachaById(id string) (model.GachaResponse, error)
	PlayGacha(model.Gacha) (model.GachaResponse, error)
}

type gachaUsecase struct {
	gr repository.IGachaRepository
	gv validator.IGachaValidator
}

func NewGachaUsecase(gr repository.IGachaRepository, gv validator.IGachaValidator) IGachaUsecase {
	return &gachaUsecase{gr, gv}
}

func (gu *gachaUsecase) GetGachaById(id string) (model.GachaResponse, error) {

	gachaResponse := model.GachaResponse{}

	return gachaResponse, nil
	// if err := gu.gr.GetGachaById(id)

	// }

}

func (gu *gachaUsecase) PlayGacha(gacha model.Gacha) (model.GachaResponse, error) {

	if err := gu.gv.GachaPlayValidator(gacha); err != nil {
		return model.GachaResponse{}, err
	}

	if err := gu.gr.CreateGacha(&gacha); err != nil {
		return model.GachaResponse{}, err
	}

	// if err := gu.gr.GetGachaById(id)
	gachaResponse := model.GachaResponse{}

	gachaResponse.Title = "hogehoge-ta"

	return gachaResponse, nil
}

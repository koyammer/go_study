package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IGachaValidator interface {
	GachaPlayValidator(gacha model.Gacha) error
}

type GachaValidator struct{}

func NewGachaValidator() IGachaValidator {
	return &GachaValidator{}
}

func (gv *GachaValidator) GachaPlayValidator(gacha model.Gacha) error {
	return validation.ValidateStruct(&gacha,
		validation.Field(
			&gacha.UserId,
			validation.Required.Error("userid is required"),
		),
	)
}

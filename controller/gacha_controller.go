package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type IGachaController interface {
	GetGachaById(c echo.Context) error
	PlayGacha(c echo.Context) error
}

type gachaController struct {
	gu usecase.IGachaUsecase
}

func NewGachaController(gu usecase.IGachaUsecase) IGachaController {
	return &gachaController{gu}
}

type QueryParams struct {
	id string `query:"id" validate:"required,max=10"`
}

func (gu *gachaController) GetGachaById(c echo.Context) error {
	params := new(QueryParams)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := c.QueryParam("id")

	gachaRes, err := gu.gu.GetGachaById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, gachaRes)

}

func (gu *gachaController) PlayGacha(c echo.Context) error {
	gacha := model.Gacha{}

	if err := c.Bind(&gacha); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	gachaRes, err := gu.gu.PlayGacha(gacha)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, gachaRes)

}

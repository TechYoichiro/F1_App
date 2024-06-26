package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TechYoichiro/F1_App/usecase"
)

type RaceController struct {
	raceUsecase usecase.RaceUsecase
}

func NewRaceController(ru usecase.RaceUsecase) *RaceController {
	return &RaceController{
		raceUsecase: ru,
	}
}

func (rc *RaceController) GetRaceData(c *gin.Context) {
	race, err := rc.raceUsecase.GetRaceData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, race)
}

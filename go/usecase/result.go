package usecase

import (
	"errors"
	"fmt"

	"github.com/TechYoichiro/F1_App/domain"
	"github.com/TechYoichiro/F1_App/interfaces"
)

type RaceUsecase interface {
	GetRaceData() (*domain.Race, error)
	PrintRaceData() error
}

type raceUsecase struct {
	raceRepo interfaces.RaceRepository
}

func NewRaceUsecase(rr interfaces.RaceRepository) RaceUsecase {
	return &raceUsecase{
		raceRepo: rr,
	}
}

func (ru *raceUsecase) GetRaceData() (*domain.Race, error) {
	url := "http://ergast.com/api/f1/current/last/results.json"
	apiResponse, err := ru.raceRepo.FetchRaceData(url)
	if err != nil {
		return nil, err
	}

	if len(apiResponse.MRData.RaceTable.Races) == 0 {
		return nil, errors.New("no race data found")
	}

	return &apiResponse.MRData.RaceTable.Races[0], nil
}

func (ru *raceUsecase) PrintRaceData() error {
	race, err := ru.GetRaceData()
	if err != nil {
		return err
	}

	fmt.Printf("シーズン: %s\nラウンド: %s\nレース名: %s\n", race.Season, race.Round, race.RaceName)
	for _, result := range race.Results {
		fmt.Printf("POS: %s, DRIVER: %s %s, TEAM: %s, TIME: %s, POINT: %s\n",
			result.Position, result.Driver.GivenName, result.Driver.FamilyName, result.Constructor.Name, result.Time.Time, result.Points)
	}

	return nil
}

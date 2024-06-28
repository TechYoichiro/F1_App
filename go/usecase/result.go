package usecase

import (
	"errors"
	"fmt"

	"github.com/TechYoichiro/F1_App/domain"
	"github.com/TechYoichiro/F1_App/interfaces"
)

// レースデータの取得と表示を行うユースケースインターフェース
type RaceUsecase interface {
	GetRaceData() (*domain.Race, error)
	PrintRaceData() error
}

type raceUsecase struct {
	raceRepo interfaces.RaceRepository
}

// 新しいRaceUsecaseを初期化
func NewRaceUsecase(rr interfaces.RaceRepository) RaceUsecase {
	return &raceUsecase{
		raceRepo: rr,
	}
}

// レースデータを取得し、Race構造体を返す
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

// レースデータを取得し、標準出力に表示する
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

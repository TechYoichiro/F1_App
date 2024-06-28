package interfaces

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/TechYoichiro/F1_App/domain"
)

// レースデータを取得するインターフェース
type RaceRepository interface {
	FetchRaceData(url string) (*domain.APIResponse, error)
}

type raceRepository struct{}

// 新しいRaceRepositoryを初期化
func NewRaceRepository() RaceRepository {
	return &raceRepository{}
}

// 指定されたURLからレースデータを取得し、APIResponse構造体にアンマーシャルする
func (r *raceRepository) FetchRaceData(url string) (*domain.APIResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch race data")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse domain.APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

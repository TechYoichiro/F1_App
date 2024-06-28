package domain

// APIレスポンスの構造体定義
type APIResponse struct {
	MRData struct {
		RaceTable struct {
			Season string `json:"season"`
			Round  string `json:"round"`
			Races  []Race `json:"Races"`
		} `json:"RaceTable"`
	} `json:"MRData"`
}

// レースデータの構造体定義
type Race struct {
	Season   string   `json:"season"`
	Round    string   `json:"round"`
	RaceName string   `json:"raceName"`
	Results  []Result `json:"Results"`
}

// レース結果の構造体定義
type Result struct {
	Position    string      `json:"position"`
	Points      string      `json:"points"`
	Driver      Driver      `json:"Driver"`
	Constructor Constructor `json:"Constructor"`
	Time        Time        `json:"Time"`
}

// ドライバーの構造体定義
type Driver struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

// コンストラクタの構造体定義
type Constructor struct {
	Name string `json:"name"`
}

// タイムの構造体定義
type Time struct {
	Time string `json:"time"`
}

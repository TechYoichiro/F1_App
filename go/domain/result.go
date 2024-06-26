package domain

type APIResponse struct {
	MRData struct {
		RaceTable struct {
			Season string `json:"season"`
			Round  string `json:"round"`
			Races  []Race `json:"Races"`
		} `json:"RaceTable"`
	} `json:"MRData"`
}

type Race struct {
	Season   string   `json:"season"`
	Round    string   `json:"round"`
	RaceName string   `json:"raceName"`
	Results  []Result `json:"Results"`
}

type Result struct {
	Position    string      `json:"position"`
	Points      string      `json:"points"`
	Driver      Driver      `json:"Driver"`
	Constructor Constructor `json:"Constructor"`
	Time        Time        `json:"Time"`
}

type Driver struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type Constructor struct {
	Name string `json:"name"`
}

type Time struct {
	Time string `json:"time"`
}

package response

type (
	Coordinate struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
	}
	History struct {
		Dist        float64      `json:"dist"`
		Hour        string       `json:"hour"`
		Coordinates []Coordinate `json:"coordinates"`
	}

	Histories struct {
		Histories []History `json:"histories"`
	}
)

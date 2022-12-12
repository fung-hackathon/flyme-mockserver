package response

type (
	Friends struct {
		Friends []User `json:"friends"`
	}
	Requests struct {
		Requests []User `json:"friends"`
	}
)

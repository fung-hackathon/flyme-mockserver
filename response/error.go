package response

type (
	Error struct {
		Code        uint   `json:"code"`
		Description string `json:"description"`
	}
)

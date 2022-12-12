package response

type (
	User struct {
		UserID   string `json:"userID"`
		UserName string `json:"userName"`
		Icon     string `json:"icon"`
	}
	Login struct {
		UserID string `json:"fun"`
		Passwd string `json:"passwd"`
	}
	Token struct {
		Token string `json:"token"`
	}
)

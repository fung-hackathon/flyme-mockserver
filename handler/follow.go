package handler

import (
	"flyme-mockserver/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MockFriends(c echo.Context) error {
	return c.JSON(http.StatusOK, &response.Friends{
		Friends: []response.User{
			{
				UserID:   "fun",
				UserName: "公立はこだて未来大学",
				Icon:     "string",
			},
		},
	})
}

func MockRequests(c echo.Context) error {
	return c.JSON(http.StatusOK, &response.Requests{
		Requests: []response.User{
			{
				UserID:   "fun",
				UserName: "公立はこだて未来大学",
				Icon:     "string",
			},
		},
	})
}

package handler

import (
	"flyme-mockserver/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MockUser(c echo.Context) error {
	return c.JSON(http.StatusOK, &response.User{
		UserID:   "fun",
		UserName: "公立はこだて未来大学",
		Icon:     "string",
	})
}

func MockUserWithParam(c echo.Context) error {
	id := c.Param("userID")
	return c.JSON(http.StatusOK, &response.User{
		UserID:   id,
		UserName: "公立はこだて未来大学",
		Icon:     "string",
	})
}

func MockToken(c echo.Context) error {
	id := c.Param("userID")
	return c.JSON(http.StatusOK, &response.User{
		UserID:   id,
		UserName: "公立はこだて未来大学",
		Icon:     "string",
	})
}

package handler

import (
	"flime-mockserver/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func MockHistory(c echo.Context) error {
	return c.JSON(http.StatusOK, &response.History{
		Dist: 33.4,
		Hour: "3:34",
		Coordinates: []response.Coordinate{
			{
				Latitude:  41.841806,
				Longitude: 140.766944,
			},
		},
	})
}

func MockHistories(c echo.Context) error {
	if n, err := strconv.ParseUint(c.QueryParam("number"), 10, 32); err != nil {
		return c.JSON(http.StatusOK, &response.Error{
			Code:        40,
			Description: "out of range",
		})
	} else if n == 0 {
		return c.JSON(http.StatusOK,
			&response.Error{
				Code:        40,
				Description: "out of range",
			})
	} else {
		number := uint(n)
		var histories []response.History
		for i := 0; uint(i) < number; i += 1 {
			histories = append(histories, response.History{
				Dist: 33.4,
				Hour: "3:34",
				Coordinates: []response.Coordinate{
					{
						Latitude:  41.841806,
						Longitude: 140.766944,
					},
				},
			})
		}
		return c.JSON(http.StatusOK, response.Histories{
			Histories: histories,
		})
	}
}

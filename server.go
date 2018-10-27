package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type LocationResponse struct {
	ID        string `json:"id"`
	RequestID string `json:"request_id"`
}

func lookupIdentifier(c echo.Context) error {
	id := c.Param("id")
	reqID := c.Response().Header().Get(echo.HeaderXRequestID)
	// summary := c.QueryParam("summary")
	return c.JSON(http.StatusOK, &LocationResponse{
		ID:        id,
		RequestID: reqID,
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.GET("/api/location/:id", lookupIdentifier)
	e.Logger.Fatal(e.Start(":1323"))
}

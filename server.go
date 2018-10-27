package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	geoip2 "github.com/oschwald/geoip2-golang"
	maxminddb "github.com/oschwald/maxminddb-golang"
)

// LocationResponse is the JSON response structure containing the geographic
// location information
type LocationResponse struct {
	ID          string `json:"id"`
	RequestID   string `json:"request_id"`
	Copyright   string `json:"copyright_notice"`
	geoip2.City `json:"geolocation"`
}

type LocationSummaryResponse struct {
	ID        string `json:"id"`
	RequestID string `json:"request_id"`
	Copyright string `json:"copyright_notice"`
	City      string `json:"geolocation"`
}

func geolocate(ipAddress string) geoip2.City {
	db, err := geoip2.Open(os.Getenv("CITY_MMDB_PATH"))
	if err != nil {
		panic(fmt.Sprintf("Unable to open geoIP2 database: %s", err))
	}
	defer db.Close()
	ip := net.ParseIP(ipAddress)

	record, err := db.City(ip)
	if err != nil {
		panic(fmt.Sprintf("Unable to lookup city: %s", err))
	}
	return *record
}

func geolocateSummary(ipAddress string) string {
	db, err := maxminddb.Open(os.Getenv("CITY_MMDB_PATH"))
	if err != nil {
		panic(fmt.Sprintf("Unable to open geoIP2 database: %s", err))
	}
	defer db.Close()
	ip := net.ParseIP(ipAddress)

	var record struct {
		Country struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"country"`
	}

	err = db.Lookup(ip, &record)
	if err != nil {
		panic(fmt.Sprintf("Unable to lookup: %s", err))
	}
	return record.Country.Names["en"]
}

func handleIdentifierLookup(c echo.Context) error {
	id := c.Param("id")
	reqID := c.Response().Header().Get(echo.HeaderXRequestID)

	if c.QueryParam("summary") != "" {
		return c.JSONPretty(http.StatusOK, &LocationSummaryResponse{
			ID:        id,
			City:      geolocateSummary(id),
			RequestID: reqID,
			Copyright: "This product includes GeoLite2 data created by MaxMind, available from http://www.maxmind.com",
		}, "  ")
	}
	return c.JSONPretty(http.StatusOK, &LocationResponse{
		ID:        id,
		City:      geolocate(id),
		RequestID: reqID,
		Copyright: "This product includes GeoLite2 data created by MaxMind, available from http://www.maxmind.com",
	}, "  ")
}

func loadEnv() {
	env := os.Getenv("GEOLOCATOR_ENV")
	if "" == env {
		env = "development"
		if os.Setenv("GEOLOCATOR_ENV", env) != nil {
			log.Fatal("Unable to set GEOLOCATOR_ENV to sane default")
		}
	}
	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env
}

func validateEnvVars() {
	if os.Getenv("CITY_MMDB_PATH") == "" {
		log.Fatal("CITY_MMDB_PATH not set. Did you follow the README?")
	}
}

func main() {
	loadEnv()
	validateEnvVars()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.GET("/api/location/:id", handleIdentifierLookup)
	e.Logger.Fatal(e.Start(":1323"))
}

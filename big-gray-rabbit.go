package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
)

const (
	API_URL = "http://transitdata.cityofmadison.com/TripUpdate/TripUpdates.json"
)

var KNOWN_STOPS = map[string]string{
	"0181": "S Park at Cedar (NB)",
	"1660": "E Washington at N Paterson (WB)",
}

func main() {
	client := &http.Client{}

	request, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	feed := gtfs.FeedMessage{}
	err = json.Unmarshal(body, &feed)
	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range feed.Entity {
		tripUpdate := entity.GetTripUpdate()
		trip := tripUpdate.GetTrip()
		tripId := trip.GetTripId()
		fmt.Printf("Trip ID: %s\n", tripId)

		stopTimeUpdates := tripUpdate.GetStopTimeUpdate()
		for _, stopTimeUpdate := range stopTimeUpdates {
			stopId := stopTimeUpdate.GetStopId()
			arrival := stopTimeUpdate.GetArrival()
			if arrival != nil {
				arrivalTime := arrival.GetTime()
				arrivalTimeFormatted := time.Unix(arrivalTime, 0).Format("2006-01-02 15:04:05")
				fmt.Printf("Stop ID: %s, Arrival Time: %s\n", stopId, arrivalTimeFormatted)
			}
		}
	}

	for stopID, stopName := range KNOWN_STOPS {
		fmt.Printf("Stop ID: %s, Stop Name: %s\n", stopID, stopName)
	}
}

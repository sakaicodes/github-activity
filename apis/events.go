package apis

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sakaicodes/github-activity/models"
)

func GetEvents(username string) ([]byte, int) {
	// API call to GitHub to get the events for a user
	// Construct the API URL with the provided username
	//Return the response body and status code
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+username+"/events", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body, resp.StatusCode
}

func GetEventsByType(username string, eventType string) ([]byte, int) {
	// API call to GitHub to get the events for a user and filter by event type
	// Construct the API URL with the provided username
	// Return the response body and status code
	marshalled_events := []models.Event{}
	filtered_events := []models.Event{}

	resp, status := GetEvents(username)
	if status != http.StatusOK {
		return nil, status
	}
	err := json.Unmarshal(resp, &marshalled_events)
	if err != nil {
		panic(err)
	}

	for _, event := range marshalled_events {
		// Condition to check if the event type matches the specified event type or if no event type is specified (empty string)
		if eventType == "" || event.Type == eventType {
			filtered_events = append(filtered_events, event)
		}
	}
	data, err := json.Marshal(filtered_events)

	if err != nil {
		panic(err)
	}
	return data, http.StatusOK

}

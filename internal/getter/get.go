package getter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/timdeklijn/druktezoeker/internal/crowdedness"
	_ "github.com/timdeklijn/druktezoeker/internal/log"
)

// getDate returns the current date in the format YYYY-MM-DD.
func getDate() string {
	currentDate := time.Now()
	return currentDate.Format("2006-01-02")
}

// buildURL builds a URL for the NS API.
func buildURL(host, station, date string) string {
	return fmt.Sprintf("%s/sigma/crowdedness/station/%s/%s", host, station, date)
}

// buildRequest builds a request for the NS API.
func buildRequest(apiKey, url string, trainNumbers []string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)

	q := req.URL.Query()
	for _, trainNumber := range trainNumbers {
		q.Add("train_numbers", trainNumber)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

// Crowdedness returns the crowdedness of a station for a given date.
func Crowdedness(config *Config, station string, trainNumbers []string) (*crowdedness.Responses, error) {
	var responses crowdedness.Responses

	date := getDate()
	logrus.Infof("Searching for trains on '%s' from '%s' for %v", date, station, trainNumbers)

	// Prep request
	url := buildURL(config.Host, station, date)
	client := &http.Client{}
	req, err := buildRequest(config.ApiKey, url, trainNumbers)
	if err != nil {
		return nil, err
	}

	// Do request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	// Parse and return response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &responses); err != nil {
		return nil, err
	}

	logrus.Infof("%+v", responses[0])
	return &responses, nil
}
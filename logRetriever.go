package convoso

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"time"
)

// LogResponse The payload that is sent back when requesting call logs from
// Convoso
type LogResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Offset     int         `json:"offset"`
		Limit      int         `json:"limit"`
		TotalFound interface{} `json:"total_found"`
		Entries    int         `json:"entries"`
		Results    []CallLog   `json:"results"`
	} `json:"data"`
}

// GetLogs return the result of requsting logs from a certain time frame,
// represented by two date structs
func GetLogs(start, end time.Time) (LogResponse, error) {
	result := LogResponse{}

	if start.After(end) {
		return result, errors.New("Starting time can not be before ending time")
	}

	data := url.Values{}
	data.Add("start_time", start.Format("2006-01-02 15:04:05"))
	data.Add("end_time", end.Format("2006-01-02 15:04:05"))
	r, err := postFormRequest("https://api.convoso.com/v1/log/retrieve", data)

	if err != nil {
		return result, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Errorf("RAW JSON: " + string(body))
		return result, err
	}

	return result, nil
}

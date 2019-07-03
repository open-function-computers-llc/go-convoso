package convoso

import (
	"net/http"
	"net/url"
	"strings"
)

func postFormRequest(route string, body url.Values) (*http.Response, error) {
	body.Add("auth_token", apiKEY)
	log.Info("POST: "+route, "POSTBODY: "+body.Encode())

	request, err := http.NewRequest("POST", route, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, _ := client.Do(request)

	return resp, nil
}

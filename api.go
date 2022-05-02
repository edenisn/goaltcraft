package altcraft

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	BaseDomain = "ru.altkraft.com"
	Version    = "v1.1"
)

type API struct {
	url   string
	Token string
}

// New creates a API client
func New(token string) *API {
	url := fmt.Sprintf("https://%s/api/%s/", BaseDomain, Version)
	return &API{
		url:   url,
		Token: token,
	}
}

// Request will make a call to the actual API
func (api *API) Request(method, path string, body map[string]interface{}) ([]byte, error) {
	client := &http.Client{}
	requestURL := fmt.Sprintf("%s/%s", api.url, path)

	var bodyBytes io.Reader
	var err error
	var data []byte
	if body != nil {
		body["token"] = api.Token
		data, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyBytes = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, requestURL, bodyBytes)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", data)
	}

	return data, nil
}

package goaltcraft

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
func (api *API) Request(method, path string, body, response interface{}) (*http.Response, error) {
	client := &http.Client{}
	requestURL := fmt.Sprintf("%s/%s", api.url, path)

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, requestURL, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return resp, nil
}

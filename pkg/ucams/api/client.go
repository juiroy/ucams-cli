package ucams_api

import (
	"fmt"
	"io"
	"net/http"
)

type (
	Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UcamsApi struct {
		host       string
		httpClient *http.Client
	}
)

func Create(host string) *UcamsApi {
	httpClient := &http.Client{}

	return &UcamsApi{
		host:       host,
		httpClient: httpClient,
	}
}

func (c *UcamsApi) do(method, endpoint string, params map[string]string, body io.Reader) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

package ucams_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	authURL = "api/v0/auth/"
)

const (
	defaultTokenTTL = 3600
)

type (
	AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		TTL      int    `json:"ttl"`
	}

	AuthResponse struct {
		Token    string `json:"token"`
		Username string `json:"username"`
	}
)

func (c *UcamsApi) Auth(username string, password string) (resp AuthResponse, err error) {
	jsonBody, err := json.Marshal(&AuthRequest{
		Username: username,
		Password: password,
		TTL:      defaultTokenTTL,
	})
	if err != nil {
		return resp, err
	}

	bodyReader := bytes.NewReader(jsonBody)

	res, err := c.do(http.MethodPost, authURL, nil, bodyReader)
	if err != nil {
		return resp, err
	}

	if res.StatusCode != 200 {
		return resp, fmt.Errorf(res.Status)
	}

	defer res.Body.Close()

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}
	log.Printf("Code: %v, Raw response: %v", res.StatusCode, string(responseBytes))
	if err = json.Unmarshal(responseBytes, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}

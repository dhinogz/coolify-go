package coolifygo

import (
	"fmt"
	"net/http"
)

type Healthcheck struct {
	Status string `json:"status"`
}

func (c *Client) GetHealthcheck() (*Healthcheck, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/healthcheck", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res := Healthcheck{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

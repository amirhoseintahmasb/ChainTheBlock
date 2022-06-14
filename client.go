package go1inch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	inchURL = "https://api.1inch.io/v4.0"
)

var (
	network = map[string]string{
		"eth":      "1",
		"bsc":      "56",
		"matic":    "137",
		"optimism": "10",
		"arbitrum": "42161",
	}
)

func NewClient() *Client {
	return &Client{
		Http: &http.Client{},
	}
}

func setQueryParam(endpoint *string, params []map[string]interface{}) {
	var first = true
	for _, param := range params {
		for i := range param {
			if first {
				*endpoint = fmt.Sprintf("%s?%s=%v", *endpoint, i, param[i])
				first = false
			} else {
				*endpoint = fmt.Sprintf("%s&%s=%v", *endpoint, i, param[i])
			}
		}
	}
}

func (c *Client) doRequest(ctx context.Context, net, endpoint, method string, expRes interface{}, reqData interface{}, opts ...map[string]interface{}) (int, error) {
	n, ok := network[net]
	if !ok {
		return 0, errors.New("invalid network")
	}
	callURL := fmt.Sprintf("%s%s%s", inchURL, n, endpoint)

	var dataReq []byte
	var err error

	if reqData != nil {
		dataReq, err = json.Marshal(reqData)
		if err != nil {
			return 0, err
		}
	}

	if len(opts) > 0 && len(opts[0]) > 0 {
		setQueryParam(&callURL, opts)
	}
	req, err := http.NewRequestWithContext(ctx, method, callURL, bytes.NewBuffer(dataReq))
	if err != nil {
		return 0, err
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	switch resp.StatusCode {
	case 200:
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				return 0, err
			}
		}
		return resp.StatusCode, nil

	default:
		return resp.StatusCode, fmt.Errorf("%s", body)
	}
}

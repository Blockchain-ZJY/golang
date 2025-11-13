package client

import (
	"context"
	"encoding/json"
	"fmt"
	"golang/microservice/types"
	"net/http"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		httErr := map[string]interface{}{}
		if err := json.NewDecoder(resp.Body).Decode(&httErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code: %d, error: %v", resp.StatusCode, httErr["error"])
	}

	priceResponse := &types.PriceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(priceResponse); err != nil {
		return nil, err
	}
	return priceResponse, nil
}

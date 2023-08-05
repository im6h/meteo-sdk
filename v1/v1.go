package meteo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/im6h/meteo-sdk/types"
)

var baseURL = "https://api.open-meteo.com/v1/forecast"

type Client struct {
	httpClient *http.Client
}

func NewClient(
	client *http.Client,
) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		httpClient: client,
	}
}

func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	client.Timeout = time.Second * 10
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// API Listing

// GetHourlyData
func (c *Client) GetHourlyData(lat, long string, props ...string) (*types.Hourly, error) {
	params := url.Values{}

	params.Add("latitude", lat)
	params.Add("longitude", long)

	if len(props) > 0 {
		params.Add("hourly", strings.Join(props, ","))
		params.Add("models", "best_match")
	} else {
		params.Add("hourly", "temperature_2m")
	}

	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	var data *types.Hourly
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetDailyData
func (c *Client) GetDailyData(lat, long, timezone string, props ...string) (*types.Daily, error) {
	params := url.Values{}

	params.Add("latitude", lat)
	params.Add("longitude", long)
	params.Add("timezone", timezone)

	if len(props) > 0 {
		params.Add("dailly", strings.Join(props, ","))
		params.Add("models", "best_match")
	} else {
		params.Add("daily", "weathercode")
		params.Add("timezone", "GMT")
	}

	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	var data *types.Daily
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

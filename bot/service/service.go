package service

import (
	"errors"
	"fmt"
	"net/http"
)

func GetStock(stock string) (*http.Response, error) {
	uri := fmt.Sprintf("https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv", stock)
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid response")
	}

	return resp, nil
}

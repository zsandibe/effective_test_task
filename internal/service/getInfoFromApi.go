package service

import (
	"effective/config"
	"effective/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getInfoFromApi(regNum string, cfg config.Config) (domain.Car, error) {
	body, err := getResponseBody(cfg.Api.Url, regNum)
	if err != nil {
		return domain.Car{}, err
	}

	var response domain.Car

	if err := json.Unmarshal(body, &response); err != nil {
		return domain.Car{}, fmt.Errorf("problems with unmarshalling response: %v", err)
	}
	return response, nil
}

func getResponseBody(url, name string) ([]byte, error) {
	fullUrl := fmt.Sprintf("%s/info?regNum=%s", url, name)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to requesting: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return body, nil
}

package api

import (
	"currency-converter-cli/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetSupportedCodes(url string) ([][]string, error) {
	var code models.Code
	if err := fetchAndDecode(url+"/codes", &code); err != nil {
		return nil, err
	}
	return code.SupportedCodes, nil
}

func GetPairConversion(url, from, to string) (*models.Conversion, error) {
	var c models.Conversion

	if err := fetchAndDecode(fmt.Sprintf("%s/pair/%s/%s", url, from, to), &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func doGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func fetchAndDecode[T any](url string, target *T) error {
	res, err := doGetRequest(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(target)
}

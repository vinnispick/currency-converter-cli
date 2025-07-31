package api

import (
	"currency-converter-cli/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetSupportedCodes(url string) [][]string {
	req, err := http.NewRequest("GET", url+"/codes", nil)
	if err != nil {
		log.Fatalf("failed to generate new request %v", err)
	}
	req.Header.Set("User-Agent", "TestAgent2")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http for supported codes failed with : %v", err)
	}
	defer res.Body.Close()
	var code models.Code

	err = json.NewDecoder(res.Body).Decode(&code)
	if err != nil {
		log.Fatalf("failed to decode JSON response: %v", err)
	}

	return code.SupportedCodes
}

func GetPairConversion(url, from, to string) models.Conversion {
	reqUrl := url + "/pair/" + from + "/" + to
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Fatalf("failed to generate new request %v", err)
	}
	req.Header.Set("User-Agent", "TestAgent2")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http for currency pair failed with : %v", err)
	}
	defer res.Body.Close()
	var c models.Conversion

	err = json.NewDecoder(res.Body).Decode(&c)
	if err != nil {
		log.Fatalf("failed to decode JSON response: %v", err)
	}

	return c

}

package chuckNorrisClient

import (
	"encoding/json"
	"io"
	"net/http"
)

type ChuckNorrisDTO struct {
	IconUrl string `json:"iconUrl"`
	Id      string `json:"id"`
	Value   string `json:"value"`
	Url     string `json:"url"`
}

func getEndpoint() string {
	return "https://api.chucknorris.io/jokes"
}

func GetFact() (ChuckNorrisDTO, error) {
	resp, err := http.Get(getEndpoint() + "/random")

	if err != nil {
		return ChuckNorrisDTO{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return ChuckNorrisDTO{}, err
	}

	var result ChuckNorrisDTO

	err = json.Unmarshal(body, &result)

	if err != nil {
		return ChuckNorrisDTO{}, err
	}

	return result, nil
}

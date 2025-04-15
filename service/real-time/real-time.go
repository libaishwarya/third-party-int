package realtime

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"third-party-int/service"
)

type RealTimeJoke struct{}

func (r *RealTimeJoke) GetJokes() (service.Jokes, error) {
	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	// if err != nil {
	// 	return service.Jokes{}, err
	// }
	// defer resp.Body.Close()

	// var jokeResponse struct {
	// 	Types     string `json:"types"`
	// 	Setup     string `json:"setup"`
	// 	Punchline string `json:"punchline"`
	// }

	// if err := json.NewDecoder(resp.Body).Decode(&jokeResponse); err != nil {
	// 	return service.Jokes{}, err
	// }
	// return service.Jokes{
	// 	Types:     jokeResponse.Types,
	// 	Setup:     jokeResponse.Setup,
	// 	Punchline: jokeResponse.Punchline,
	// }, nil
	if err != nil {
		return service.Jokes{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return service.Jokes{}, err

	}

	var jokeResponse struct {
		Types     string `json:"type"`
		Setup     string `json:"setup"`
		Punchline string `json:"punchline"`
	}

	if err := json.Unmarshal(body, &jokeResponse); err != nil {
		return service.Jokes{}, err
	}

	return service.Jokes{
		Types:     jokeResponse.Types,
		Setup:     jokeResponse.Setup,
		Punchline: jokeResponse.Punchline,
	}, nil
}

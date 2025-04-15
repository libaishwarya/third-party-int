package store

type Joke struct {
	Types     string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type JokesStore interface {
	StoreJoke(jokes *Joke) error
}

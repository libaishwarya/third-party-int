package service

type Jokes struct {
	Types     string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type JokeService interface {
	GetJokes() (Jokes, error)
}

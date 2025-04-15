package mock

import (
	"fmt"
	"third-party-int/service"
)

type MockJokes struct {
	Fail bool
}

func (m *MockJokes) GetJokes() (service.Jokes, error) {
	if m.Fail {
		return service.Jokes{}, fmt.Errorf("error fetching jokes")
	}
	return service.Jokes{
		Types:     "general",
		Setup:     "What do you get hanging from Apple trees?",
		Punchline: "Sore arms.",
	}, nil
}

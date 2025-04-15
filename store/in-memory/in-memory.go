package inmemory

import "third-party-int/store"

type InmemoryJokeStore struct {
	jokes []store.Joke
}

func NewInmemoryJokeStore() *InmemoryJokeStore {
	return &InmemoryJokeStore{jokes: []store.Joke{}}
}

func (s *InmemoryJokeStore) StoreJoke(joke *store.Joke) error {
	s.jokes = append(s.jokes, *joke)
	return nil
}

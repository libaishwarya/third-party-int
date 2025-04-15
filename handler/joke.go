package handler

import (
	"net/http"
	"third-party-int/service"
	"third-party-int/store"

	"github.com/gin-gonic/gin"
)

type JokeHandler struct {
	storeH store.JokesStore
	jokes  service.JokeService
}

func NewJokeHandler(store store.JokesStore, jokehandler service.JokeService) *JokeHandler {
	return &JokeHandler{storeH: store, jokes: jokehandler}
}

func (h *JokeHandler) StoreJokes(c *gin.Context) {
	jokes, err := h.jokes.GetJokes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch jokes"})
	}
	joke := &store.Joke{
		Types:     jokes.Types,
		Setup:     jokes.Setup,
		Punchline: jokes.Punchline,
	}
	if err := h.storeH.StoreJoke(joke); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store jokes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Joke stored successfully"})
}

package routes

import (
	"third-party-int/handler"
	"third-party-int/service"
	"third-party-int/store"

	"github.com/gin-gonic/gin"
)

func SetupRouter(jokeStore store.JokesStore, jokeHandler service.JokeService) *gin.Engine {
	r := gin.Default()

	h := handler.NewJokeHandler(jokeStore, jokeHandler)
	AttachRoutes(h, r)
	return r
}

func AttachRoutes(h *handler.JokeHandler, r *gin.Engine) {
	r.GET("/jokes", h.StoreJokes)

}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/service/joke"
)

// Joke is a struct that holds dependencies the needed by the handler.
// REF: https://github.com/gin-gonic/gin/issues/932#issuecomment-306242400
type Joke struct {
	Service joke.Service
}

// JokeHandler retrieves a list of available jokes
func (j Joke) JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	val, err := j.Service.GetJokes(c)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, val)
}

// LikeJoke increments the likes of a particular joke Item
func (j *Joke) LikeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	jokeid, err := strconv.Atoi(c.Param("jokeID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	val, err := j.Service.LikeJokes(c, strconv.Itoa(jokeid))
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, val)
}

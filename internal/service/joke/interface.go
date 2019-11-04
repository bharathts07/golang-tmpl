package joke

import "context"

// Joke contains information about a single Joke
type Joke struct {
	ID    string `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Text  string `json:"joke" binding:"required"`
}

type Service interface {
	GetJokes(ctx context.Context) ([]Joke, error)
	LikeJokes(ctx context.Context, jokeID string) ([]Joke, error)
}

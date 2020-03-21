package joke

import (
	"context"
	"errors"
)

// jokes is a list of jokes along with id and number of likes present by default
var defaultJokes = []Joke{
	{"1", 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	{"2", 0, "What do you call a fake noodle? An Impasta."},
	{"3", 0, "How many apples grow on a tree? All of them."},
	{"4", 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	{"5", 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	{"6", 0, "Why did the coffee file a police report? It got mugged."},
	{"7", 0, "How does a penguin build it's house? Igloos it together."},
}

type jokeImpl struct {
	JokesCollection []Joke
}

func (s *jokeImpl) GetJokes(ctx context.Context) ([]Joke, error) {
	jokes := make([]Joke, len(s.JokesCollection))
	copy(jokes, s.JokesCollection)
	return jokes, nil
}

func (s *jokeImpl) LikeJokes(ctx context.Context, jokeID string) ([]Joke, error) {
	for i := 0; i < len(s.JokesCollection); i++ {
		if s.JokesCollection[i].ID == jokeID {
			s.JokesCollection[i].Likes++
			return s.JokesCollection, nil
		}
	}
	return nil, errors.New("joke not found")
}

func New() Service {
	startJokes := make([]Joke, len(defaultJokes))
	copy(startJokes, defaultJokes)
	return &jokeImpl{
		JokesCollection: startJokes,
	}
}

package fakedb

import (
	"github.com/bharathts07/pokke/database"
)

type implementation struct {
}

func New() database.Client {
	return &implementation{}
}

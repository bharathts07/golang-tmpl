package fakedb

import (
	"context"
	"errors"

	"github.com/bharathts07/pokke/internal/models"
)

func (db implementation) GetComponents(ctx context.Context,
	view string, components []models.ComponentType) ([]models.Component, error) {
	switch view {
	case "welcome":
		components := []models.Component{
			{
				Type:models.Body,
				Data:models.Welcome{
					Message: "ようこそう",
					Home:    "https://0.0.0.0:8080/home",
				},
			},
		}
		return components,nil
	case "home":
	case "contact":
	}
	return nil, errors.New("invalid view")
}

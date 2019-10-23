package fakedb

import (
	"context"
	"errors"

	"github.com/bharathts07/pokke/internal/models"
)

func (db implementation) GetLayout(ctx context.Context, view string) (*models.Layout, error) {
	switch view {
	case "welcome":
		layout := models.Layout{
			Alignment:    "center",
			Order:        models.OrderMap{
				Blocks:[]int{
					1,
				},
			},
			ComponentIDs: []models.ComponentType{
				models.Body,
			},
		}
		return &layout,nil
	case "home":
		layout := models.Layout{
			Alignment:"center",
			Order:models.OrderMap{
				Blocks: []int{1,1,1,1},
			},
			ComponentIDs: []models.ComponentType{
				models.Banner,
				models.MenuBar,
				models.Body,
				models.Footer,
			},
		}
		return &layout,nil
	case "contact":
	}
	return nil, errors.New("invalid view")
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/internal/database/blog/model"
	"github.com/bharathts07/pokke/internal/service/blog"
)

type Blog struct {
	Service blog.Service
}

func (h Blog) GetAll(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// parse the limit
	blogTitle, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		blogTitle = 10
	}

	val, err := h.Service.BlogPreview(c, blogTitle)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, val)
}

func (h Blog) GetBlog(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// parse the request
	blogTitle := c.Param("title")
	val, err := h.Service.ReadBlog(c, blogTitle)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, val)
}

func (h Blog) PostBlog(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// parse the request
	var blogItem model.BlogItem
	err := c.Bind(&blogItem)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}
	// save in the service
	err = h.Service.SaveBlog(c, &blogItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, blogItem.Title)
}

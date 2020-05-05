package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/upbit/ploy_illusts/model"
)

var (
	errInvalidID = errors.New("Invalid ID")
)

// GetIllust Endpoint
func GetIllust(c *gin.Context) {
	var retErr error
	defer func() {
		if retErr != nil {
			c.JSON(400, gin.H{
				"error":   true,
				"message": retErr.Error(),
			})
		}
	}()

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		retErr = err
		return
	}

	illust, err := model.GetIllust(ID)
	if err != nil {
		retErr = err
		return
	}

	c.JSON(200, gin.H{"illust": illust})
}

// GetIllusts Endpoint
func GetIllusts(c *gin.Context) {
	var retErr error
	defer func() {
		if retErr != nil {
			c.JSON(400, gin.H{
				"error":   true,
				"message": retErr.Error(),
			})
		}
	}()

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}
	// 排序方式
	var sortFields []string
	sort := c.Query("sort")
	switch sort {
	case "hot":
		sortFields = []string{"-total_bookmarks", "-total_view"}
	default: // default
		sortFields = []string{"-_id"}
	}

	illusts, err := model.GetIllusts(page, size, sortFields)
	if err != nil {
		retErr = err
		return
	}

	c.JSON(200, gin.H{
		"page":    page,
		"size":    size,
		"sort":    sort,
		"illusts": illusts,
	})
}

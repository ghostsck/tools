package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

var PageSize = 1

func GetGinPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * PageSize
	}
	return result
}

package gintool

import "github.com/gin-gonic/gin"

import "go-learn/internal/pkg/str"

// GetPaginationParam 获取分页参数
func GetPaginationParam(c *gin.Context) (page, size int) {
	return str.AtoiDefault(c.Query("page"), 0), str.AtoiDefault(c.Query("size"), 20)
}

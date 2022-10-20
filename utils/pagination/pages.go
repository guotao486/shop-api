/*
 * @Author: GG
 * @Date: 2022-10-20 15:11:00
 * @LastEditTime: 2022-10-20 15:44:08
 * @LastEditors: GG
 * @Description:pagination 分页工具类
 * @FilePath: \shop-api\utils\pagination\pages.go
 *
 */
package pagination

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	// 默认页数
	DefaultPageSize = 100
	// 最大页数
	MaxPageSize = 1000
	// 查询参数名称
	PageVar = "page"
	// 页数查询参数名称
	PageSizeVar = "pageSize"
)

// 分页结构体
type Pages struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"pageSize"`
	// 总页数
	PageCount int `json:"pageCount"`
	// 最大数量
	TotalCount int         `json:"totalCount"`
	Items      interface{} `json:"items"`
}

// 实例化分页结构体
func New(page, pageSize, total int) *Pages {
	// 每页数量
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}

	// 每页最大数量
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	pageCount := -1
	if total >= 0 {
		// 最大页码
		pageCount = (total + pageSize - 1) / pageSize
		if page > pageCount {
			page = pageCount
		}
	}
	// 页码处理
	if page <= 0 {
		page = 1
	}

	return &Pages{
		Page:       page,
		PageSize:   pageSize,
		PageCount:  pageCount,
		TotalCount: total,
	}
}

// 根据http请求实例化分页结构体
func NewFromRequest(req *http.Request, count int) *Pages {
	page := ParseInt(req.URL.Query().Get(PageVar), 1)
	pageSize := ParseInt(req.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

// 根据gin请求实例化分页结构体
func NewFromGinRequest(g *gin.Context, count int) *Pages {
	page := ParseInt(g.Query(PageVar), 1)
	pageSize := ParseInt(g.Query(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

// string转换int
func ParseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// offset 分页位移
func (p *Pages) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// limit
func (p *Pages) Limit() int {
	return p.PageSize
}

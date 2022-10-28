/*
 * @Author: GG
 * @Date: 2022-10-28 15:31:59
 * @LastEditTime: 2022-10-28 16:05:22
 * @LastEditors: GG
 * @Description: category controller
 * @FilePath: \shop-api\api\category\controller.go
 *
 */
package category

import (
	"fmt"
	"shopping/domain/category"
	"shopping/utils/api_helper"
	"shopping/utils/pagination"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	categoryService *category.Service
}

/**
 * @description: 实例化
 * @param {*category.Service} categoryService
 * @return {*}
 */
func NewCategoryController(categoryService *category.Service) *Controller {
	return &Controller{
		categoryService: categoryService,
	}
}

// CreateCategory godoc
// @Summary 根据给定的参数创建分类
// @Tags Category
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param CreateCategoryRequest body CreateCategoryRequest true "category information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /category [post]
func (c *Controller) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest

	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	newCategory := category.NewCategory(req.Name, req.Desc)

	if err := c.categoryService.Create(newCategory); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, api_helper.Response{
		Message: "create category success!",
	})
}

// BulkCreateCategory godoc
// @Summary 根据给定的csv文件，批量创建分类
// @Tags Category
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param   file formData file true  "file contains category information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /category/upload [post]
func (c *Controller) BulkCreateCategory(g *gin.Context) {
	fileHeader, err := g.FormFile("file")
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	count, err := c.categoryService.BulkCreate(fileHeader)
	if err != nil {
		api_helper.HandleError(g, err)
	}

	api_helper.HandleSuccess(g, api_helper.Response{
		Message: fmt.Sprintf("'%s' uploaded! '%d' new categories created", fileHeader.Filename, count),
	})
}

// GetCategories godoc
// @Summary 获得分类列表
// @Tags Category
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} pagination.Pages
// @Router /category [get]
func (c *Controller) GetCategorys(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = c.categoryService.GetAll(page)
	api_helper.HandleSuccess(g, page)
}

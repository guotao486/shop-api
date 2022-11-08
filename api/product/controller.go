/*
 * @Author: GG
 * @Date: 2022-10-28 16:23:35
 * @LastEditTime: 2022-11-07 15:03:43
 * @LastEditors: GG
 * @Description: product controller
 * @FilePath: \shop-api\api\product\controller.go
 *
 */
package product

import (
	"errors"
	"shopping/domain/product"
	"shopping/utils/api_helper"
	"shopping/utils/pagination"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	productService *product.Service
}

/**
 * @description: 实例化
 * @param {product.Service} productService
 * @return {*}
 */
func NewProductController(productService product.Service) *Controller {
	return &Controller{
		productService: &productService,
	}
}

// GetProducts godoc
// @Summary 获得商品列表（分页）
// @Tags Product
// @Accept json
// @Produce json
// @Param qt query string false "Search text to find matched sku numbers and names"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} pagination.Pages
// @Router /product [get]
func (c *Controller) GetProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	queryText := g.Query("qt")
	c.productService.GetAllBySearch(queryText, page)
	api_helper.HandleSuccess(g, page)
}

// CreateProduct godoc
// @Summary 创建商品
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization  header    string  true  "Authentication header"
// @Param CreateProductRequest body CreateProductRequest true "product information"
// @Success 200 {object} CreateProductResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /product [post]
func (c *Controller) CreateProduct(g *gin.Context) {
	var req CreateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.CreateProduct(req.Name, req.Desc, req.Count, req.Price, req.CategoryID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, CreateProductResponse{
		Message: "create product success!",
	})
}

// UpdateProduct godoc
// @Summary 更新商品
// @Tags Product
// @Accpet json
// @Produce json
// @Param Authorization  header    string  true  "Authentication header"
// @Param UpdateProductRequest body UpdateProductRequest true "product information"
// @Success 200 {object} UpdateProductResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /product [patch]
func (c *Controller) UpdateProduct(g *gin.Context) {
	var req UpdateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	if req.SKU == "" {
		api_helper.HandleError(g, errors.New("请选择商品"))
		return
	}
	var currentProduct *product.Product
	currentProduct, err := c.productService.FindProductBySku(req.SKU)
	if err != nil {
		api_helper.HandleError(g, product.ErrProductNotFound)
		return
	}

	err = c.productService.UpdateProduct(req.ToProduct(currentProduct))
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, UpdateProductResponse{
		Message: "Product Update Success!",
	})
}

// DeleteProduct godoc
// @Summary 删除商品根据sku
// @Tags Product
// @Accept json
// @Produce json
// @Param DeleteProductRequest body DeleteProductRequest true "sku of product"
// @Param Authorization header    string  true  "Authentication header"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /product [delete]
func (c *Controller) DeleteProduct(g *gin.Context) {
	var req DeleteProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.DeleteProduct(req.SKU)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, api_helper.Response{
		Message: "Delete Product Success!",
	})
}

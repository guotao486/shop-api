/*
 * @Author: GG
 * @Date: 2022-10-29 10:28:29
 * @LastEditTime: 2022-11-05 16:59:57
 * @LastEditors: GG
 * @Description: cart controller
 * @FilePath: \shop-api\api\cart\controller.go
 *
 */
package cart

import (
	"shopping/domain/cart"
	"shopping/utils/api_helper"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cartService *cart.Service
}

/**
 * @description: 实例化
 * @param {*cart.Service} cartService
 * @return {*}
 */
func NewCartController(cartService *cart.Service) *Controller {
	return &Controller{
		cartService: cartService,
	}
}

// AddItem godoc
// @Summary 添加Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} ItemCartResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart/item [post]
func (c *Controller) AddItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	userId := api_helper.GetUserId(g)
	err := c.cartService.AddItem(userId, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, ItemCartResponse{
		Message: "Add Cart Item Success!",
	})
}

// UpdateItem godoc
// @Summary 更新Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart/item [patch]
func (c *Controller) UpdateItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	userId := api_helper.GetUserId(g)
	err := c.cartService.UpdateItem(userId, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, ItemCartResponse{
		Message: "Update Cart Item Success!",
	})
}

// GetCart godoc
// @Summary 获得购物车商品列表
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {array} cart.Item
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart [get]
func (c *Controller) GetItems(g *gin.Context) {
	userId := api_helper.GetUserId(g)
	result, err := c.cartService.GetCartItem(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, result)
}

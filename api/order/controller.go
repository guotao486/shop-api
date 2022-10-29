/*
 * @Author: GG
 * @Date: 2022-10-29 13:43:09
 * @LastEditTime: 2022-10-29 14:14:56
 * @LastEditors: GG
 * @Description: order controller
 * @FilePath: \shop-api\api\order\controller.go
 *
 */
package order

import (
	"shopping/domain/order"
	"shopping/utils/api_helper"
	"shopping/utils/pagination"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	orderService *order.Service
}

/**
 * @description: 实例化
 * @param {*order.Service} orderService
 * @return {*}
 */
func NewOrderController(orderService *order.Service) *Controller {
	return &Controller{
		orderService: orderService,
	}
}

// CompleteOrder godoc
// @Summary 完成订单
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /order [post]
func (c *Controller) CompleteOrder(g *gin.Context) {
	userId := api_helper.GetUserId(g)

	err := c.orderService.CompleteOrder(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, CompleteOrderResponse{
		Message: "Order Success!",
	})
}

// CancelOrder godoc
// @Summary 取消订单
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param CancelOrderRequest body CancelOrderRequest true "order information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /order [delete]
func (c *Controller) CancelOrder(g *gin.Context) {
	var req CancelOrderRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	userId := api_helper.GetUserId(g)
	err := c.orderService.CancelOrder(userId, req.OrderId)

	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, CancelOrderResponse{
		Message: "Cancel Order Success!",
	})
}

// GetOrders godoc
// @Summary 获得订单列表
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} pagination.Pages
// @Router /order [get]
func (c *Controller) GetAll(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	userId := api_helper.GetUserId(g)
	page = c.orderService.GetAll(page, userId)
	api_helper.HandleSuccess(g, page)
}

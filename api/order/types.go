/*
 * @Author: GG
 * @Date: 2022-10-29 13:41:11
 * @LastEditTime: 2022-10-29 13:50:25
 * @LastEditors: GG
 * @Description: order types
 * @FilePath: \shop-api\api\order\types.go
 *
 */
package order

type CompleteOrderRequest struct {
}
type CompleteOrderResponse struct {
	Message string `json:"message"`
}
type CancelOrderRequest struct {
	OrderId uint `json:"orderId"`
}
type CancelOrderResponse struct {
	Message string `json:"message"`
}

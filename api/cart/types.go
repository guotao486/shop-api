/*
 * @Author: GG
 * @Date: 2022-10-29 09:55:37
 * @LastEditTime: 2022-10-29 10:28:17
 * @LastEditors: GG
 * @Description: cart types
 * @FilePath: \shop-api\api\cart\types.go
 *
 */
package cart

type ItemCartRequest struct {
	SKU   string `json:"sku"`
	Count int    `json:"count"`
}

type ItemCartResponse struct {
	Message string `json:"message"`
}

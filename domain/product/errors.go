/*
 * @Author: GG
 * @Date: 2022-10-21 10:29:47
 * @LastEditTime: 2022-10-21 10:31:43
 * @LastEditors: GG
 * @Description:product errors
 * @FilePath: \shop-api\domain\product\errors.go
 *
 */
package product

import "errors"

var (
	ErrProductNotFound        = errors.New("商品没有找到")
	ErrProductStockIsNoEnough = errors.New("商品库存不足")
)

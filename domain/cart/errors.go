/*
 * @Author: GG
 * @Date: 2022-10-24 11:29:34
 * @LastEditTime: 2022-10-24 17:11:43
 * @LastEditors: GG
 * @Description:cart errors
 * @FilePath: \shop-api\domain\cart\errors.go
 *
 */
package cart

import "errors"

var (
	ErrItemAlreadyExistInCart = errors.New("商品已经存在")
	ErrCountInvalid           = errors.New("数量不正确")
	ErrItemNotFound           = errors.New("找不到该宝贝~")
)

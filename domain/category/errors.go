/*
 * @Author: GG
 * @Date: 2022-10-18 16:31:42
 * @LastEditTime: 2022-10-18 16:32:46
 * @LastEditors: GG
 * @Description:category errors
 * @FilePath: \shop-api\domain\category\errors.go
 *
 */
package category

import "errors"

var (
	ErrCategoryExistWithName = errors.New("商品分类已存在")
)

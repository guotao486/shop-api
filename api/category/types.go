/*
 * @Author: GG
 * @Date: 2022-10-28 14:52:57
 * @LastEditTime: 2022-10-28 14:54:34
 * @LastEditors: GG
 * @Description: category types
 * @FilePath: \shop-api\api\category\types.go
 *
 */
package category

type CreateCategoryRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CreateCategoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

/*
 * @Author: GG
 * @Date: 2022-10-28 16:07:54
 * @LastEditTime: 2022-10-28 16:21:52
 * @LastEditors: GG
 * @Description: product controller types
 * @FilePath: \shop-api\api\product\types.go
 *
 */
package product

import "shopping/domain/product"

type CreateProductRequest struct {
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}

type CreateProductResponse struct {
	Message string `json:"message"`
}

type DeleteProductRequest struct {
	SKU string `json:"sku"`
}

type UpdateProductRequest struct {
	SKU        string  `json:"sku"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}

type UpdateProductResponse struct {
	Message string `json:"message"`
}

/**
 * @description: 类型转换，UpdateProductRequest to Product
 * @return {*}
 */
func (p *UpdateProductRequest) ToProduct() *product.Product {
	return product.NewProduct(p.Name, p.Desc, p.Count, p.Price, p.CategoryID)
}

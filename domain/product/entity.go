/*
 * @Author: GG
 * @Date: 2022-10-21 10:20:27
 * @LastEditTime: 2022-10-21 10:31:54
 * @LastEditors: GG
 * @Description:product model
 * @FilePath: \shop-api\domain\product\entity.go
 *
 */
package product

import (
	"shopping/domain/category"

	"gorm.io/gorm"
)

// 商品结构体
type Product struct {
	gorm.Model
	Name       string
	SKU        string
	Desc       string
	StockCount int
	Price      float32
	CategoryID uint
	Category   category.Category `json:"-"`
	IsDelete   bool
}

/**
 * @name: 实例化商品结构体
 * @description:
 * @param {string} name
 * @param {string} desc
 * @param {int} stockCount
 * @param {float32} price
 * @param {uint} cid
 * @return {*}
 */
func NewProduct(name string, desc string, stockCount int, price float32, cid uint) *Product {

	return &Product{
		Name:       name,
		Desc:       desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: cid,
		IsDelete:   false,
	}
}

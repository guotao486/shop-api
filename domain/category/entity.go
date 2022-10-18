/*
 * @Author: GG
 * @Date: 2022-10-18 16:28:03
 * @LastEditTime: 2022-10-18 16:45:05
 * @LastEditors: GG
 * @Description:category model
 * @FilePath: \shop-api\domain\category\entity.go
 *
 */
package category

import "gorm.io/gorm"

// 商品分类结构体
type Category struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Desc     string
	IsActive bool
}

/**
 * @name:
 * @description:
 * @param {string} name
 * @param {string} desc
 * @return {*}
 */
func NewCategory(name string, desc string) *Category {
	return &Category{
		Name:     name,
		Desc:     desc,
		IsActive: true,
	}
}

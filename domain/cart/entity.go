/*
 * @Author: GG
 * @Date: 2022-10-24 10:02:48
 * @LastEditTime: 2022-10-24 11:28:54
 * @LastEditors: GG
 * @Description:cart model cart_item model
 * @FilePath: \shop-api\domain\cart\entity.go
 *
 */
package cart

import (
	"shopping/domain/product"
	"shopping/domain/user"

	"gorm.io/gorm"
)

/**
 * @description: 购物车结构体
 * @return {*}
 */
type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `gorm:"foreignKey:UserID;"`
}

/**
 * @description: 实例化Cart
 * @param {uint} userId
 * @return {*}
 */
func NewCart(userId uint) *Cart {
	return &Cart{
		UserID: userId,
	}
}

// item 结构体
type Item struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID;"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

/**
 * @description: 实例化CartItem
 * @param {uint} productId
 * @param {uint} cartId
 * @param {int} count
 * @return {*}
 */
func NewCartItem(productId uint, cartId uint, count int) *Item {
	return &Item{
		ProductID: productId,
		CartID:    cartId,
		Count:     count,
	}
}

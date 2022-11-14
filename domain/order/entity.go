/*
 * @Author: GG
 * @Date: 2022-10-25 09:54:40
 * @LastEditTime: 2022-11-14 15:23:19
 * @LastEditors: GG
 * @Description: Order and OrderItem model
 * @FilePath: \shop-api\domain\order\entity.go
 *
 */
package order

import (
	"shopping/domain/product"
	"shopping/domain/user"

	"gorm.io/gorm"
)

// 订单结构体
type Order struct {
	gorm.Model
	UserID     uint
	User       user.User   `gorm:"foreignKey:UserID" json:"-"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	TotalPrice float32
	IsCanceled bool
}

// 订单项结构体
type OrderItem struct {
	gorm.Model
	Product    product.Product `gorm:"foreignKey:ProductID"`
	ProductID  uint
	Count      int
	OrderID    uint
	IsCanceled bool
	Price      float32
}

/**
 * @description: 实例化订单
 * @param {uint} uid
 * @param {[]OrderItem} items
 * @return {*}
 */
func NewOrder(uid uint, items []OrderItem) *Order {
	var totalPrice float32 = 0.0
	for _, item := range items {
		totalPrice += item.Price
	}
	return &Order{
		UserID:     uid,
		OrderItems: items,
		TotalPrice: totalPrice,
		IsCanceled: false,
	}
}

/**
 * @description: 实例化订单选项
 * @param {int} count
 * @param {uint} pid
 * @return {*}
 */
func NewOrderItem(count int, pid uint, price float32) *OrderItem {
	return &OrderItem{
		ProductID:  pid,
		Count:      count,
		IsCanceled: false,
		Price:      price,
	}
}

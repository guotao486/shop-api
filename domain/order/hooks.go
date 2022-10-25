/*
 * @Author: GG
 * @Date: 2022-10-25 10:11:17
 * @LastEditTime: 2022-10-25 11:56:05
 * @LastEditors: GG
 * @Description:
 * @FilePath: \shop-api\domain\order\hooks.go
 *
 */
package order

import (
	"shopping/domain/cart"
	"shopping/domain/product"

	"gorm.io/gorm"
)

/**
 * @description: 创建订单之前，检查购物车
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	var currentCart cart.Cart
	// 检查购物车
	if err := tx.Where("user_id = ?", order.UserID).First(&currentCart).Error; err != nil {
		return err
	}
	return nil
}

/**
 * @description: 创建订单之后，删除购物车和商品
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (order *Order) AfterCreate(tx *gorm.DB) (err error) {
	var currentCart cart.Cart
	// 检查购物车
	if err := tx.Where("user_id = ?", order.UserID).First(&currentCart).Error; err != nil {
		return err
	}
	// 删除购物车商品
	if err := tx.Where("cart_id = ?", currentCart.ID).Unscoped().Delete(&cart.Item{}).Error; err != nil {
		return err
	}
	// 删除购物车
	if err := tx.Unscoped().Delete(&currentCart).Error; err != nil {
		return err
	}
	return nil
}

/**
 * @description: 保存订单商品之前，更新产品库存
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (orderItem *OrderItem) BeforeSave(tx *gorm.DB) (err error) {
	var currentProduct product.Product
	var currentOrderItem OrderItem

	// 查询商品是否存在
	if err := tx.Where("id = ?", orderItem.ProductID).First(currentProduct).Error; err != nil {
		return err
	}

	// 订单商品数量
	reservedStockCount := currentOrderItem.Count

	// 要扣除的库存
	newStockCount := currentProduct.StockCount - reservedStockCount
	if newStockCount < 0 {
		return ErrNotEnoughStock
	}

	// 乐观锁，更新商品库存
	if err := tx.Model(&currentProduct).Where("stockCount = ?", currentProduct.StockCount).Update("stockCount", newStockCount).Error; err != nil {
		return err
	}

	if orderItem.Count == 0 {
		err := tx.Unscoped().Delete(&orderItem).Error
		return err
	}

	return nil
}

/**
 * @description: 订单取消之后，返还库存
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	if order.IsCanceled {
		var orderItems []OrderItem
		if err := tx.Where("order_id = ?", order.ID).First(&orderItems).Error; err != nil {
			return err
		}

		for _, item := range orderItems {
			var currentProduct product.Product

			if err := tx.Where("id = ?", item.ProductID).First(&currentProduct).Error; err != nil {
				return err
			}

			newStockCount := currentProduct.StockCount + item.Count
			if err := tx.Model(&currentProduct).Where("stockCount = ?", currentProduct.StockCount).Update("stockCount = ?", newStockCount).Error; err != nil {
				return err
			}
			if err := tx.Model(&item).Update("IsCanceled", true).Error; err != nil {
				return err
			}
		}
	}
	return
}

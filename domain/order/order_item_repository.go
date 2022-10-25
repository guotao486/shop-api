/*
 * @Author: GG
 * @Date: 2022-10-25 15:28:37
 * @LastEditTime: 2022-10-25 15:40:44
 * @LastEditors: GG
 * @Description: orderItem repository
 * @FilePath: \shop-api\domain\order\order_item_repository.go
 *
 */
package order

import (
	"log"

	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

/**
 * @description: 实例化
 * @param {*gorm.DB} db
 * @return {*}
 */
func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{
		db: db,
	}
}

/**
 * @description: 创建表
 * @return {*}
 */
func (r *OrderItemRepository) Migration() {

	err := r.db.AutoMigrate(&OrderItem{})
	if err != nil {
		log.Println(err)
	}
}

/**
 * @description: 创建orderItem
 * @param {*OrderItem} item
 * @return {*}
 */
func (r *OrderItemRepository) Create(item *OrderItem) error {
	result := r.db.Create(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * @description: 更新orderItem
 * @param {OrderItem} newItem
 * @return {*}
 */
func (r *OrderItemRepository) Update(newItem OrderItem) error {
	result := r.db.Save(&newItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

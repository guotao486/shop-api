/*
 * @Author: GG
 * @Date: 2022-10-25 15:05:23
 * @LastEditTime: 2022-11-14 16:11:20
 * @LastEditors: GG
 * @Description:order repository
 * @FilePath: \shop-api\domain\order\order_repository.go
 *
 */
package order

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

/**
 * @description: 实例化
 * @param {*gorm.DB} db
 * @return {*}
 */
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

/**
 * @description: 创建表
 * @return {*}
 */
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Order{})
	if err != nil {
		log.Println(err)
	}
}

/**
 * @description: 订单创建
 * @param {*Order} order
 * @return {*}
 */
func (r *Repository) Create(order *Order) error {
	result := r.db.Create(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * @description: 订单更新
 * @param {Order} newOrder
 * @return {*}
 */
func (r *Repository) Update(newOrder *Order) error {
	result := r.db.Updates(&newOrder)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * @description: 根据id查询订单
 * @param {uint} oid
 * @return {*}
 */
func (r *Repository) FindByOrderId(oid uint) (*Order, error) {
	var currentOrder Order
	if err := r.db.Where("IsCanceled = ?", false).Where("id = ?", oid).First(&currentOrder).Error; err != nil {
		return nil, err
	}
	return &currentOrder, nil
}

/**
 * @description: 获取所有订单，分页列表
 * @param {*} pageIndex
 * @param {int} pageSize
 * @param {uint} uid
 * @return {*}
 */
func (r *Repository) GetAll(pageIndex, pageSize int, uid uint) ([]Order, int) {
	var orders []Order
	var count int64

	// 查询订单
	r.db.Where("IsCanceled = ?", false).Where("UserID = ?", uid).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Count(&count)

	for i, order := range orders {
		// 订单项数据
		r.db.Where("OrderID = ?", order.ID).Find(&orders[i].OrderItems)

		for j, item := range orders[i].OrderItems {
			// 订单项商品数据
			r.db.Where("ID = ?", item.ProductID).First(&orders[i].OrderItems[j].Product)
		}
	}

	return orders, int(count)
}

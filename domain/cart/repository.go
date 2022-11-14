/*
 * @Author: GG
 * @Date: 2022-10-24 15:31:47
 * @LastEditTime: 2022-11-14 15:16:04
 * @LastEditors: GG
 * @Description: cart repository
 * @FilePath: \shop-api\domain\cart\repository.go
 *
 */
package cart

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// 实例化
func NewCartRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

/**
 * @description: 创建表
 * @return {*}
 */
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Cart{})
	if err != nil {
		log.Print(err)
	}
}

/**
 * @description: 根据userId查找购物车，找不到则根据userId新建一个
 * @param {uint} userId
 * @return {*}
 */
func (r *Repository) CartFindOrCreateByUserId(userId uint) (*Cart, error) {
	var cart *Cart
	err := r.db.Where(Cart{UserID: userId}).Attrs(NewCart(userId)).FirstOrCreate(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

/**
 * @description: 根据userId查找购物车
 * @param {uint} userId
 * @return {*}
 */
func (r *Repository) CartFindByUserId(userId uint) (*Cart, error) {
	var cart *Cart
	err := r.db.Where(Cart{UserID: userId}).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

/**
 * @description: 更新购物车
 * @param {Cart} cart
 * @return {*}
 */
func (r *Repository) CartUpdate(cart Cart) error {
	result := r.db.Updates(cart)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//////////////// Cart Item  ////////////////
type ItemRepository struct {
	db *gorm.DB
}

/**
 * @description: 实例化
 * @param {*gorm.DB} db
 * @return {*}
 */
func NewCartItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

/**
 * @description: 创建Item表
 * @return {*}
 */
func (r *ItemRepository) Migration() {
	err := r.db.AutoMigrate(&Item{})
	if err != nil {
		log.Print(err)
	}
}

/**
 * @description: 创建购物车商品
 * @param {Cart} cart
 * @return {*}
 */
func (r *ItemRepository) Create(ci *Item) error {
	result := r.db.Create(&ci)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

/**
 * @description: 更新购物车商品
 * @param {Item} ci
 * @return {*}
 */
func (r *ItemRepository) Update(ci *Item) error {
	result := r.db.Updates(&ci)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * @description: 根据商品id和购物车id查找购物车商品
 * @param {uint} pid
 * @param {uint} cid
 * @return {*}
 */
func (r *ItemRepository) FindById(pid uint, cid uint) (*Item, error) {
	var item Item
	err := r.db.Where(Item{ProductID: pid, CartID: cid}).First(&item).Error
	if err != nil {
		return nil, ErrItemNotFound
	}
	return &item, nil
}

/**
 * @description: 返回购物车所有item
 * @param {uint} cartId
 * @return {*}
 */
func (r *ItemRepository) GetItems(cartId uint) ([]Item, error) {
	var items []Item
	err := r.db.Where(Item{CartID: cartId}).Find(&items).Error
	if err != nil {
		return nil, err
	}

	// 查找关联的商品属性
	for i, item := range items {
		err := r.db.Model(item).Association("Product").Find(&items[i].Product)
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}

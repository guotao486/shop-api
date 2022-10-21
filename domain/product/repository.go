/*
 * @Author: GG
 * @Date: 2022-10-21 10:42:35
 * @LastEditTime: 2022-10-21 16:12:30
 * @LastEditors: GG
 * @Description:product dao repository
 * @FilePath: \shop-api\domain\product\repository.go
 *
 */
package product

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

/**
 * @name:实例化
 * @description:
 * @param {*gorm.DB} db
 * @return {*}
 */
func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

/**
 * @name: 创建表
 * @description:
 * @return {*}
 */
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Product{})
	if err != nil {
		log.Println(err)
	}
}

/**
 * @description: 创建商品
 * @param {*Product} product
 * @return {*}
 */
func (r *Repository) Create(product *Product) error {
	result := r.db.Create(product)
	return result.Error
}

/**
 * @description: 根据sku查询商品
 * @param {string} sku
 * @return {*}
 */
func (r *Repository) FindBySku(sku string) (*Product, error) {
	var product *Product
	err := r.db.Where("IsDelete = ?", 0).Where(Product{SKU: sku}).First(&product).Error
	if err != nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}

/**
 * @description: 更新商品
 * @param {Product} updateProduct
 * @return {*}
 */
func (r *Repository) Update(updateProduct Product) error {
	saveProduct, err := r.FindBySku(updateProduct.SKU)
	if err != nil {
		return err
	}
	err = r.db.Model(&saveProduct).Updates(updateProduct).Error
	return err
}

/**
 * @description: 检索所有商品，分页列表
 * @param {string} search
 * @param {*} pageIndex
 * @param {int} pageSize
 * @return {*}
 */
func (r *Repository) GetAllBySearch(search string, pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64

	query := r.db.Where("IsDelete =?", 0)
	if search != "" {
		query = query.Where("name like ? or SKU like ?", search, search)
	}

	query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

/**
 * @description: 根据sku删除
 * @param {string} sku
 * @return {*}
 */
func (r *Repository) Delete(sku string) error {
	product, err := r.FindBySku(sku)
	if err != nil {
		return err
	}

	product.IsDelete = true

	err = r.db.Save(product).Error
	return err
}

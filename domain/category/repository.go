/*
 * @Author: GG
 * @Date: 2022-10-18 16:46:51
 * @LastEditTime: 2022-10-20 17:15:08
 * @LastEditors: GG
 * @Description: category dao repository
 * @FilePath: \shop-api\domain\category\repository.go
 *
 */
package category

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// 创建商品分类
func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成商品分类
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
}

/**
 * @name: 插入测试数据
 * @description:
 * @return {*}
 */

func (r *Repository) InsertSampleDate() {
	categorys := []Category{
		{Name: "cat1", Desc: "category 1"},
		{Name: "cat2", Desc: "category 2"},
	}
	for _, c := range categorys {
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}

/**
 * @name: 创建分类
 * @description:
 * @param {*Category} c
 * @return {*}
 */
func (r *Repository) Create(c *Category) error {
	result := r.db.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * @name: 删除分类
 * @description:
 * @param {*Category} c
 * @return {*}
 */
func (r *Repository) Delete(c *Category) error {
	result := r.db.Delete(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * @name: 更新分类
 * @description:
 * @param {*Category} c
 * @return {*}
 */
func (r *Repository) Update(c *Category) error {
	result := r.db.Save(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * @name: 根据名称查询分类
 * @description:
 * @param {string} name
 * @return {*}
 */
func (r *Repository) GetByName(name string) []Category {
	var categorys []Category
	r.db.Where("name = ?", name).Find(&categorys)
	return categorys
}

/**
 * @name:分页查询分类
 * @description:
 * @param {int} pageIndex
 * @param {int} pageSize
 * @return {*}
 */
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categorys []Category
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categorys).Count(&count)
	return categorys, int(count)
}

/**
 * @name: 批量创建分类
 * @description:
 * @param {*[]Category} categorys
 * @return {*}
 */
func (r *Repository) BulkCreate(categorys []*Category) (int, error) {
	var count int64
	err := r.db.Create(&categorys).Count(&count).Error
	return int(count), err
}

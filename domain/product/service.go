/*
 * @Author: GG
 * @Date: 2022-10-21 16:14:10
 * @LastEditTime: 2022-11-07 15:15:17
 * @LastEditors: GG
 * @Description:product service
 * @FilePath: \shop-api\domain\product\service.go
 *
 */
package product

import "shopping/utils/pagination"

type Service struct {
	r Repository
}

/**
 * @description: 实例化服务类
 * @param {Repository} r
 * @return {*}
 */
func NewProductService(r Repository) *Service {
	r.Migration()
	return &Service{
		r: r,
	}
}

/**
 * @description: 检索所有商品，分页列表
 * @param {string} search
 * @param {*pagination.Pages} page
 * @return {*}
 */
func (s *Service) GetAllBySearch(search string, page *pagination.Pages) *pagination.Pages {
	products, count := s.r.GetAllBySearch(search, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

/**
 * @description: 创建商品
 * @param {string} name
 * @param {string} desc
 * @param {int} count
 * @param {float32} price
 * @param {uint} cid
 * @return {*}
 */
func (s *Service) CreateProduct(name string, desc string, count int, price float32, cid uint) error {
	product := NewProduct(name, desc, count, price, cid)
	err := s.r.Create(product)
	return err
}

/**
 * @description: 更新商品
 * @param {*Product} product
 * @return {*}
 */
func (s *Service) UpdateProduct(product *Product) error {
	err := s.r.Update(*product)
	return err
}

/**
 * @description: 删除商品
 * @param {string} sku
 * @return {*}
 */
func (s *Service) DeleteProduct(sku string) error {
	err := s.r.Delete(sku)
	return err
}

/**
 * @description: 根据sku查找商品
 * @param {string} sku
 * @return {*}
 */
func (s *Service) FindProductBySku(sku string) (*Product, error) {
	return s.r.FindBySku(sku)
}

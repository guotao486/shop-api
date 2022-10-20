/*
 * @Author: GG
 * @Date: 2022-10-20 16:42:56
 * @LastEditTime: 2022-10-20 17:15:49
 * @LastEditors: GG
 * @Description:category service
 * @FilePath: \shop-api\domain\category\service.go
 *
 */
package category

import (
	"mime/multipart"
	"shopping/utils/csv_helper"
	"shopping/utils/pagination"
)

type Service struct {
	r Repository
}

// 初始化商品分类service
func NewCategoryService(r Repository) *Service {

	r.Migration()
	r.InsertSampleDate()
	return &Service{
		r: r,
	}
}

/**
 * @name: 创建商品分类
 * @description:
 * @param {*Category} category
 * @return {*}
 */
func (s Service) Create(category *Category) error {
	existCity := s.r.GetByName(category.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}

	err := s.r.Create(category)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categorys := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}

	for _, categoryVariables := range bulkCategory {
		categorys = append(categorys, NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	count, err := s.r.BulkCreate(categorys)
	if err != nil {
		return count, err
	}

	return count, nil
}

/**
 * @name: 获得商品分类分页列表
 * @description:
 * @param {*pagination.Pages} page
 * @return {*}
 */
func (s Service) GetAll(page *pagination.Pages) *pagination.Pages {
	categorys, count := s.r.GetAll(page.Page, page.PageSize)
	page.Items = categorys
	page.TotalCount = count
	return page
}

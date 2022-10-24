/*
 * @Author: GG
 * @Date: 2022-10-24 16:36:56
 * @LastEditTime: 2022-10-24 17:31:37
 * @LastEditors: GG
 * @Description: Cart service
 * @FilePath: \shop-api\domain\cart\service.go
 *
 */
package cart

import "shopping/domain/product"

type Service struct {
	cartRepository    Repository
	itemRepository    ItemRepository
	productRepository product.Repository
}

/**
 * @description: 实例化service
 * @param {Repository} cartRepository
 * @param {ItemRepository} itemRepository
 * @param {product.Repository} productRepository
 * @return {*}
 */
func NewService(cartRepository Repository, itemRepository ItemRepository, productRepository product.Repository) *Service {
	cartRepository.Migration()
	itemRepository.Migration()
	return &Service{
		cartRepository:    cartRepository,
		itemRepository:    itemRepository,
		productRepository: productRepository,
	}
}

/**
 * @description: 添加购物车
 * @param {uint} userId
 * @param {string} sku
 * @param {int} count
 * @return {*}
 */
func (s *Service) AddItem(userId uint, sku string, count int) error {
	// 商品数据
	currentProduct, err := s.productRepository.FindBySku(sku)
	if err != nil {
		return err
	}

	// 查找购物车，没有则新建
	currentCart, err := s.cartRepository.CartFindOrCreateByUserId(userId)
	if err != nil {
		return nil
	}

	// 检查购物车商品是否存在
	_, err = s.itemRepository.FindById(currentProduct.ID, currentCart.ID)
	if err == nil {
		return ErrItemAlreadyExistInCart
	}

	// 检查商品库存
	if currentProduct.StockCount < count {
		return product.ErrProductStockIsNoEnough
	}

	// 检查count值
	if count <= 0 {
		return ErrCountInvalid
	}

	// 加入购物车
	err = s.itemRepository.Create(NewCartItem(currentProduct.ID, currentCart.ID, count))
	return err
}

/**
 * @description: 更新购物车商品
 * @param {uint} userId
 * @param {string} sku
 * @param {int} count
 * @return {*}
 */
func (s *Service) UpdateItem(userId uint, sku string, count int) error {
	// 商品数据
	currentProduct, err := s.productRepository.FindBySku(sku)
	if err != nil {
		return err
	}

	// 查找购物车，没有则新建
	currentCart, err := s.cartRepository.CartFindOrCreateByUserId(userId)
	if err != nil {
		return nil
	}

	// 检查购物车商品是否存在
	currentItem, err := s.itemRepository.FindById(currentProduct.ID, currentCart.ID)
	if err != nil {
		return ErrItemNotFound
	}

	// 检查商品库存
	if currentProduct.StockCount < count {
		return product.ErrProductStockIsNoEnough
	}

	// 更新购物车商品数量，并更新数据
	currentItem.Count = count
	err = s.itemRepository.Update(currentItem)
	return err
}

/**
 * @description: 获得用户购物车所有商品
 * @param {uint} userId
 * @return {*}
 */
func (s *Service) GetCartItem(userId uint) ([]Item, error) {
	currentCart, err := s.cartRepository.CartFindByUserId(userId)
	if err != nil {
		return nil, err
	}

	items, err := s.itemRepository.GetItems(currentCart.ID)
	if err != nil {
		return nil, err
	}

	return items, nil
}

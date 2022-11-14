/*
 * @Author: GG
 * @Date: 2022-10-25 15:44:20
 * @LastEditTime: 2022-11-14 16:04:45
 * @LastEditors: GG
 * @Description: order service
 * @FilePath: \shop-api\domain\order\service.go
 *
 */
package order

import (
	"fmt"
	"shopping/domain/cart"
	"shopping/domain/product"
	"shopping/utils/pagination"
	"time"
)

var day14ToHours float64 = 336

type Service struct {
	orderRepository     Repository
	orderItemRepository OrderItemRepository
	productRepository   product.Repository
	cartRepository      cart.Repository
	cartItemRepository  cart.ItemRepository
}

/**
 * @description: 实例化
 * @return {*}
 */
func NewService(
	orderRepository Repository,
	orderItemRepository OrderItemRepository,
	productRepository product.Repository,
	cartRepository cart.Repository,
	cartItemRepository cart.ItemRepository,
) *Service {
	orderRepository.Migration()
	orderItemRepository.Migration()
	return &Service{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
		productRepository:   productRepository,
		cartRepository:      cartRepository,
		cartItemRepository:  cartItemRepository,
	}
}

/**
 * @description: 订单创建完成
 * @param {uint} userId
 * @return {*}
 */
func (s *Service) CompleteOrder(userId uint) error {
	currentCart, err := s.cartRepository.CartFindByUserId(userId)
	if err != nil {
		return err
	}
	cartItems, err := s.cartItemRepository.GetItems(currentCart.ID)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return ErrEmptyCartFound
	}

	orderItems := make([]OrderItem, 0)
	for i, citem := range cartItems {
		fmt.Printf("i: %v\n", i)
		orderItems = append(orderItems, *NewOrderItem(citem.Count, citem.ProductID, citem.Product.Price))
	}
	err = s.orderRepository.Create(NewOrder(userId, orderItems))
	return err
}

/**
 * @description: 取消订单
 * @param {*} uid
 * @param {uint} oid
 * @return {*}
 */
func (s *Service) CancelOrder(uid, oid uint) error {
	currentOrder, err := s.orderRepository.FindByOrderId(oid)
	if err != nil {
		return err
	}
	if currentOrder.UserID != uid {
		return ErrInvalidOrderID
	}

	// 是否有效期内
	if currentOrder.CreatedAt.Sub(time.Now()).Hours() > day14ToHours {
		return ErrCancelDurationPassed
	}
	currentOrder.IsCanceled = true
	err = s.orderRepository.Update(currentOrder)
	return err
}

/**
 * @description: 根据用户id获取全部订单，分页列表
 * @param {*pagination.Pages} page
 * @param {uint} uid
 * @return {*}
 */
func (s *Service) GetAll(page *pagination.Pages, uid uint) *pagination.Pages {
	orders, count := s.orderRepository.GetAll(page.Page, page.PageSize, uid)
	page.Items = orders
	page.TotalCount = count
	return page
}

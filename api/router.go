/*
 * @Author: GG
 * @Date: 2022-11-05 11:09:46
 * @LastEditTime: 2022-11-05 15:42:09
 * @LastEditors: GG
 * @Description: 路由文件
 * @FilePath: \shop-api\api\router.go
 *
 */
package api

import (
	"log"
	cartApi "shopping/api/cart"
	categoryApi "shopping/api/category"
	orderApi "shopping/api/order"
	productApi "shopping/api/product"
	userApi "shopping/api/user"
	"shopping/config"
	"shopping/domain/cart"
	"shopping/domain/category"
	"shopping/domain/order"
	"shopping/domain/product"
	"shopping/domain/user"
	"shopping/utils/database"
	"shopping/utils/middleware"

	"github.com/gin-gonic/gin"
)

type Databaces struct {
	categoryRepository  *category.Repository
	userRepository      *user.Repository
	cartRepository      *cart.Repository
	cartItemRepository  *cart.ItemRepository
	orderRepository     *order.Repository
	orderItemRepository *order.OrderItemRepository
	productRepository   *product.Repository
}

// 配置文件全局对象
var AppConfig = &config.Configuration{}

/**
 * @description: 根据配置文件创建数据库
 * @return {*}
 */
func CreateDBs() *Databaces {
	cfgFile := "./config/config.yaml"
	conf, err := config.GettAllConfigValues(cfgFile)
	if err != nil {
		log.Fatalln("读取配置文件失败. %v", err.Error())
		return nil
	}

	AppConfig = conf
	db := database.NewMysqlDB(AppConfig.DatabaseSettings.DatabaseURI)
	return &Databaces{
		categoryRepository:  category.NewCategoryRepository(db),
		userRepository:      user.NewUserRepository(db),
		cartRepository:      cart.NewCartRepository(db),
		cartItemRepository:  cart.NewCartItemRepository(db),
		orderRepository:     order.NewRepository(db),
		orderItemRepository: order.NewOrderItemRepository(db),
		productRepository:   product.NewProductRepository(db),
	}
}

/**
 * @description: 路由注册
 * @param {*gin.Engine} r
 * @return {*}
 */
func RegisterHandlers(r *gin.Engine) {
	dbs := *CreateDBs()
	RegisterCategoryHandlers(r, dbs)
	RegisterProductHandlers(r, dbs)
	RegisterCartHandlers(r, dbs)
	RegisterOrderHandlers(r, dbs)
	RegisterUserHandlers(r, dbs)
}

/**
 * @description: 注册分类控制器路由
 * @param {*gin.Engine} r
 * @param {Databaces} dbs
 * @return {*}
 */
func RegisterCategoryHandlers(r *gin.Engine, dbs Databaces) {
	categoryService := category.NewCategoryService(*dbs.categoryRepository)
	categoryController := categoryApi.NewCategoryController(categoryService)
	categoryGourp := r.Group("/category")
	// post /category
	categoryGourp.POST("", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.CreateCategory)
	// get /category
	categoryGourp.GET("", categoryController.GetCategorys)
	// post /category/upload
	categoryGourp.POST("/upload", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.BulkCreateCategory)
}

/**
 * @description: 注册用户控制器路由
 * @param {*gin.Engine} r
 * @param {Databaces} dbs
 * @return {*}
 */
func RegisterUserHandlers(r *gin.Engine, dbs Databaces) {
	userService := user.NewUserService(*dbs.userRepository)
	userController := userApi.NewUserController(userService, AppConfig)
	userGroup := r.Group("/user")
	// post /user
	userGroup.POST("", userController.CreateUser)
	// post /user/login
	userGroup.POST("/login", userController.Login)
}

/**
 * @description: 注册商品控制器
 * @param {*gin.Engine} r
 * @param {Databaces} dbs
 * @return {*}
 */
func RegisterProductHandlers(r *gin.Engine, dbs Databaces) {
	productService := product.NewProductService(*dbs.productRepository)
	productController := productApi.NewProductController(*productService)
	productGroup := r.Group("/product")

	// get /product
	productGroup.GET("", productController.GetProducts)
	// post /product
	productGroup.POST("", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.CreateProduct)
	// delete /product
	productGroup.DELETE("", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.DeleteProduct)
	// patch /product
	productGroup.PATCH("", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.UpdateProduct)
}

/**
 * @description: 注册购物控制器
 * @param {*gin.Engine} r
 * @param {Databaces} dbs
 * @return {*}
 */
func RegisterCartHandlers(r *gin.Engine, dbs Databaces) {
	cartService := cart.NewService(*dbs.cartRepository, *dbs.cartItemRepository, *dbs.productRepository)
	cartController := cartApi.NewCartController(cartService)

	cartGroup := r.Group("/cart", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	// get /cart
	cartGroup.GET("", cartController.GetItems)
	// patch /cart
	cartGroup.PATCH("", cartController.UpdateItem)
	// post /cart
	cartGroup.POST("", cartController.AddItem)
}

/**
 * @description: 注册订单控制器
 * @param {*gin.Engine} r
 * @param {Databaces} dbs
 * @return {*}
 */
func RegisterOrderHandlers(r *gin.Engine, dbs Databaces) {
	orderService := order.NewService(*dbs.orderRepository, *dbs.orderItemRepository, *dbs.productRepository, *dbs.cartRepository, *dbs.cartItemRepository)
	orderController := orderApi.NewOrderController(orderService)

	orderGroup := r.Group("/order", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	// get /order
	orderGroup.GET("", orderController.GetAll)
	// post /order
	orderGroup.POST("", orderController.CompleteOrder)
	// delete /order
	orderGroup.DELETE("", orderController.CancelOrder)
}

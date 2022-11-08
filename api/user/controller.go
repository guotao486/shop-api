/*
 * @Author: GG
 * @Date: 2022-10-26 16:36:23
 * @LastEditTime: 2022-11-07 14:47:12
 * @LastEditors: GG
 * @Description: user controller
 * @FilePath: \shop-api\api\user\controller.go
 *
 */
package user

import (
	"fmt"
	"os"
	"shopping/config"
	"shopping/domain/user"
	"shopping/utils/api_helper"
	jwtHelper "shopping/utils/jwt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService *user.Service
	appConfig   *config.Configuration
}

/**
 * @description: 实例化
 * @param {*user.Service} service
 * @param {*config.Configuration} appConfig
 * @return {*}
 */
func NewUserController(service *user.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		userService: service,
		appConfig:   appConfig,
	}
}

// CreateUser godoc
// @Summary 根据用户名和密码创建用户
// @Tags Auth
// @Accept json
// @Produce json
// @Param CreateUserRequest body CreateUserRequest true "user information"
// @Success 201 {object} CreateUserResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user [post]
func (c *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest
	// 参数绑定
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}

	newUser := user.NewUser(req.Username, req.Password, req.Password2)
	err := c.userService.CreateUser(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	api_helper.HandleSuccess(g, CreateUserResponse{
		Username: newUser.Username,
	})
}

// Login godoc
// @Summary 根据用户名和密码登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginRequest body LoginRequest true "user information"
// @Success 200 {object} LoginResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user/login [post]
func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}

	currentUser, err := c.userService.Login(req.Username, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	// 验证token
	decodeClaims := jwtHelper.VerifyToken(currentUser.Token, c.appConfig.SecretKey)
	fmt.Printf("decodeClaims: %v\n", decodeClaims)
	// token无效则重新生成token
	if decodeClaims == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp":      time.Now().Add(24 * time.Hour).Unix(),
				"isAdmin":  currentUser.IsAdmin,
			})

		token, err := jwtHelper.GenerateToken(jwtClaims, c.appConfig.SecretKey)
		currentUser.Token = token
		err = c.userService.UpdateUser(&currentUser)

		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
	}

	api_helper.HandleSuccess(g, LoginResponse{
		Username: currentUser.Username,
		UserId:   int(currentUser.ID),
		Token:    currentUser.Token,
	})
}

/**
 * @description: 验证token
 * @param {*gin.Context} g
 * @return {*}
 */
func (c *Controller) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.SecretKey)
	api_helper.HandleSuccess(g, decodedClaims)
}

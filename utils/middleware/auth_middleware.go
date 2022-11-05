/*
 * @Author: GG
 * @Date: 2022-11-05 10:20:50
 * @LastEditTime: 2022-11-05 10:38:02
 * @LastEditors: GG
 * @Description: 权限中间件
 * @FilePath: \shop-api\utils\middleware\auth_middleware.go
 *
 */
package middleware

import (
	"net/http"
	jwtHelper "shopping/utils/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 管理员权限
 * @param {string} secretKey
 * @return {*}
 */
func AuthAdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodeClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodeClaims != nil && decodeClaims.IsAdmin {
				c.Next()
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{"message": "没有访问权限"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "请重新登录"})
		}
		c.Abort()
		return
	}
}

/**
 * @description: 用户权限
 * @param {string} secretKey
 * @return {*}
 */
func AuthUserMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodeClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodeClaims != nil {
				// 需要用到的userId
				c.Set("userId", decodeClaims.UserId)
				c.Next()
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{"message": "没有访问权限"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "请重新登录"})
		}
		c.Abort()
		return
	}
}

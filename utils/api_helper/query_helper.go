/*
 * @Author: GG
 * @Date: 2022-10-25 16:35:38
 * @LastEditTime: 2022-10-25 16:39:56
 * @LastEditors: GG
 * @Description: user query helper
 * @FilePath: \shop-api\utils\api_helper\query_helper.go
 *
 */
package api_helper

import (
	"shopping/utils/pagination"

	"github.com/gin-gonic/gin"
)

var userIdText = "userId"

func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}

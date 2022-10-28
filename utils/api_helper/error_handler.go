/*
 * @Author: GG
 * @Date: 2022-10-25 16:45:15
 * @LastEditTime: 2022-10-28 16:05:39
 * @LastEditors: GG
 * @Description:error handle
 * @FilePath: \shop-api\utils\api_helper\error_handler.go
 *
 */
package api_helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(g *gin.Context, err error) {
	g.JSON(
		http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})

	g.Abort()
	return
}

func HandleSuccess(g *gin.Context, obj any) {
	g.JSON(http.StatusOK, obj)
	g.Abort()
	return
}

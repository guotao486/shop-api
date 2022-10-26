/*
 * @Author: GG
 * @Date: 2022-10-25 16:45:15
 * @LastEditTime: 2022-10-26 17:17:31
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
	g.JSON(http.StatusCreated, obj)
	g.Abort()
	return
}

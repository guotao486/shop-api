/*
 * @Author: GG
 * @Date: 2022-10-25 16:29:07
 * @LastEditTime: 2022-10-25 16:31:52
 * @LastEditors: GG
 * @Description: api type helper
 * @FilePath: \shop-api\utils\api_helper\types.go
 *
 */
package api_helper

import "errors"

// 响应结构体
type Response struct {
	Message string `json:"message"`
}

// 响应错误结构体
type ErrorResponse struct {
	Message string `json:"errosMessage"`
}

var (
	ErrInvalidBody = errors.New("请检查您的请求体")
)

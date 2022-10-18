/*
 * @Author: GG
 * @Date: 2022-10-17 16:28:04
 * @LastEditTime: 2022-10-18 16:32:59
 * @LastEditors: GG
 * @Description: user errors
 * @FilePath: \shop-api\domain\user\user.go
 *
 */
package user

import "errors"

var (
	ErrUserExistWithName = errors.New("用户名已经存在")
	ErrUserNotFound      = errors.New("用户未找到")

	ErrMismatchedPasswords = errors.New("密码不匹配")
	ErrInvalidUsername     = errors.New("无效用户名")
	ErrInvalidPassword     = errors.New("无效密码")
)

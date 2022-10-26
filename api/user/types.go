/*
 * @Author: GG
 * @Date: 2022-10-26 16:32:23
 * @LastEditTime: 2022-10-26 16:35:53
 * @LastEditors: GG
 * @Description: user types
 * @FilePath: \shop-api\api\user\types.go
 *
 */
package user

type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	UserId   int    `json:"userId"`
	Token    string `json:"token"`
}

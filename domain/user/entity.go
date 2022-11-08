/*
 * @Author: GG
 * @Date: 2022-10-17 15:59:03
 * @LastEditTime: 2022-11-07 10:33:09
 * @LastEditors: GG
 * @Description: User model
 * @FilePath: \shop-api\domain\user\entity.go
 *
 */
package user

import "gorm.io/gorm"

// 用户模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30)"`
	Password  string `gorm:"type:varchar(100)"`
	Password2 string `gorm:"-"`
	Salt      string `gorm:"type:varchar(100)"`
	Token     string `gorm:"type:varchar(500)"`
	IsDelete  bool
	IsAdmin   bool
}

// 新建用户实例
func NewUser(username, password, password2 string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Password2: password2,
		IsDelete:  false,
		IsAdmin:   false,
	}
}

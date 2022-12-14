/*
 * @Author: GG
 * @Date: 2022-10-17 16:58:47
 * @LastEditTime: 2022-10-17 17:03:53
 * @LastEditors: GG
 * @Description: user hooks
 * @FilePath: \shop-api\domain\user\hooks.go
 *
 */
package user

import (
	"shopping/utils/hash"

	"gorm.io/gorm"
)

// 保存用户之前回调，如果密码没有被加密,加密密码和salt
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Salt == "" {
		// 为salt创建一个随机字符串
		salt := hash.CreateSalt()
		// 创建hash加密密码
		hashPassword, err := hash.HashPassword(u.Password + salt)
		if err != nil {
			return nil
		}
		u.Salt = salt
		u.Password = hashPassword
	}

	return
}

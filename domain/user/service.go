/*
 * @Author: GG
 * @Date: 2022-10-18 15:58:42
 * @LastEditTime: 2022-11-07 10:37:06
 * @LastEditors: GG
 * @Description:user service
 * @FilePath: \shop-api\domain\user\service.go
 *
 */
package user

import (
	"shopping/utils/hash"
)

// 用户service结构体
type Service struct {
	r Repository
}

// 实例化user service
func NewUserService(r Repository) *Service {
	r.Migration()
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// 创建用户
func (s Service) CreateUser(user *User) error {
	// 验证两次密码是否一致
	if user.Password != user.Password2 {
		return ErrMismatchedPasswords
	}

	// 用户名已存在
	_, err := s.r.GetByName(user.Username)
	if err == nil {
		return ErrUserExistWithName
	}

	// 无效用户名
	if !ValidateUserName(user.Username) {
		return ErrInvalidUsername
	}

	// 无效密码
	if !ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}

	err = s.r.Create(user)
	return err
}

/**
 * 更新用户
**/
func (s Service) UpdateUser(user *User) error {
	return s.r.Update(user)
}

/**
 */
func (s Service) DeleteUser(user *User) error {
	return s.r.Delete(user)
}

func (s Service) Login(username string, password string) (User, error) {
	user, err := s.r.GetByName(username)
	if err != nil {
		return User{}, ErrUserNotFound
	}

	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

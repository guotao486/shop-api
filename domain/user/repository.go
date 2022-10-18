/*
 * @Author: GG
 * @Date: 2022-10-18 10:36:01
 * @LastEditTime: 2022-10-18 11:03:08
 * @LastEditors: GG
 * @Description: User dao
 * @FilePath: \shop-api\domain\user\repository.go
 *
 */
package user

import (
	"log"

	"gorm.io/gorm"
)

// Repository 结构体
type Repository struct {
	db *gorm.DB
}

// 实例化
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成表
func (r Repository) Migration() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}

// 创建用户
func (r Repository) Create(u *User) error {
	result := r.db.Create(u)
	return result.Error
}

// 根据用户名查询用户
func (r Repository) GetByName(name string) (User, error) {
	var user User
	err := r.db.Where("Username = ?", name).Where("IsDelete = ?", 0).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// 添加测试数据
func (r Repository) InsertSampleData() {
	user := NewUser("admin", "password", "password")
	user.IsAdmin = true
	//使用attrs来初始化参数，如果未找到数据则使用attrs中的数据来初始化一条
	// FirstOrCreate获取第一个匹配的记录，若没有，则根据条件初始化一个新的记录：
	//注意：attrs 必须 要和FirstOrInit 或者 FirstOrCreate 连用
	r.db.Where(User{Username: user.Username}).Attrs(User{Username: user.Username, Password: user.Password}).FirstOrCreate(&user)

	user = NewUser("user", "password", "password")
	r.db.Where(User{Username: user.Username}).Attrs(User{Username: user.Username, Password: user.Password}).FirstOrCreate(&user)
}

// 更新用户
func (r Repository) Update(u *User) error {
	result := r.db.Save(u)
	return result.Error
}

// 删除用户
func (r Repository) Delete(u *User) error {
	result := r.db.Delete(u)
	return result.Error
}

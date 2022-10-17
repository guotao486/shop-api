/*
 * @Author: GG
 * @Date: 2022-10-17 16:35:34
 * @LastEditTime: 2022-10-17 16:57:00
 * @LastEditors: GG
 * @Description: 密码加密工具类
 * @FilePath: \shop-api\utils\hash\hash.go
 *
 */
package hash

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 随机字符串
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 使用当前时间（纳秒）创建seed
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// 创建加密盐
func CreateSalt() string {
	b := make([]byte, bcrypt.MaxCost)
	for i := range b {
		// 根据seed返回非负伪随机数从charset中取字符
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// 使用bcrypt算法返回hash后密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 检查密码是否相等
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

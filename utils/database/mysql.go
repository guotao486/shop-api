/*
 * @Author: GG
 * @Date: 2022-11-05 10:48:39
 * @LastEditTime: 2022-11-05 10:58:32
 * @LastEditors: GG
 * @Description: mysql database
 * @FilePath: \shop-api\utils\database\mysql.go
 *
 */
package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/**
 * @description: 实例化mysql
 * @param {string} conString
 * @return {*}
 */
func NewMysqlDB(conString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
			NoLowerCase:   true, // 关闭大写转小写,这是关键,启动这个就不在会导致中文乱码!
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic(fmt.Sprintf("不能连接到数据库 : %s", err.Error()))
	}
	return db
}

/*
 * @Author: GG
 * @Date: 2022-10-21 10:32:48
 * @LastEditTime: 2022-10-21 11:23:29
 * @LastEditors: GG
 * @Description:product hooks
 * @FilePath: \shop-api\domain\product\hooks.go
 *
 */
package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.SKU = uuid.New().String()
	return
}

/*
 * @Author: GG
 * @Date: 2022-10-24 11:32:36
 * @LastEditTime: 2022-10-24 11:34:45
 * @LastEditors: GG
 * @Description:cart hooks
 * @FilePath: \shop-api\domain\cart\hooks.go
 *
 */
package cart

import "gorm.io/gorm"

func (item *Item) AfterUpdate(tx *gorm.DB) (err error) {

	if item.Count <= 0 {
		return tx.Unscoped().Delete(&item).Error
	}
	return
}

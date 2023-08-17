/*
功能：mysql事务
说明：
*/
package xmysql

import (
	"errors"

	"gorm.io/gorm"
)

// mysql事务
func Mysql_transaction(db *gorm.DB, transFunc func(*gorm.DB) error) (err error) {
	tx := db.Begin()
	err = errors.New("transaction begin")

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = transFunc(tx)
	return
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

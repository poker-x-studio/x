/*
功能：gorm打开mysql数据库连接池，并且配置连接池参数
说明：
*/
package xmysql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 打开gorm连接池，并且配置
func Open_gorm_config_conn_pool(source string, mysql_logger logger.Interface, option *GormOption) (*gorm.DB, func(), error) {
	//连接池
	db, close_handler, err := open_gorm(source, mysql_logger)
	if err != nil {
		return nil, nil, err
	}

	//配置连接池
	sqldb, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println("MaxOpenConnections:", sqldb.Stats().MaxOpenConnections)

	sqldb.SetConnMaxLifetime(time.Duration(option.Max_lifetime) * time.Second)
	sqldb.SetMaxOpenConns(option.Max_open_conns)
	sqldb.SetMaxIdleConns(option.Max_idle_conns)

	//fmt.Println("MaxOpenConnections:", sqldb.Stats().MaxOpenConnections)
	return db, close_handler, nil
}

// 打开gorm连接池
func open_gorm(source string, mysql_logger logger.Interface) (*gorm.DB, func(), error) {
	// charset=utf8&parseTime=True&loc=Local
	// if !strings.Contains(source, "?") {
	// 	source += "?parseTime=true"
	// } else {
	// 	source += "&parseTime=true"
	// }

	cfg := &gorm.Config{Logger: mysql_logger}
	// gorm2.0 没有关闭函数了
	db, err := gorm.Open(mysql.Open(source), cfg)
	if err != nil {
		return nil, nil, err
	}

	return db, func() {
		sql_db, err := db.DB()
		if err == nil && sql_db != nil {
			sql_db.Close()
		}
	}, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

/*
功能：gorm打开mysql数据库连接池，并且配置连接池参数
说明：
*/
package xdatabase

type GormOption struct {
	Max_lifetime   int `toml:"max_lifetime"`
	Max_open_conns int `toml:"max_open_conns"`
	Max_idle_conns int `toml:"max_idle_conns"`
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

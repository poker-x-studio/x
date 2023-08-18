/*
功能：redis 列表
说明：

Redis列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）
*/
package xredis

import (
	"github.com/gomodule/redigo/redis"
)

// Function:将一个或多个值插入到列表头部
// Description:LPUSH KEY_NAME VALUE1.. VALUEN
// Return:执行 LPUSH 命令后，列表的长度。
func LPush(rp *redis.Pool, key, value1 string) (int, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("LPUSH", key, value1)
	return redis.Int(reply, err)
}

// Function:Llen 命令用于返回列表的长度
// Description:LLEN KEY_NAME
// Return:列表的长度。
func LLen(rp *redis.Pool, key string) (int, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("LLEN", key)
	return redis.Int(reply, err)
}

// Function:Lrange 返回列表中指定区间内的元素，区间以偏移量 START 和 END 指定。
// Description:LRANGE KEY_NAME START END
// Return:一个列表，包含指定区间内的元素。
func LRange(rp *redis.Pool, key string, start, end int) ([]string, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("LRANGE", key, start, end)
	return redis.Strings(reply, err)
}

// Function:一个列表，包含指定区间内的元素。
// Description:LSET KEY_NAME INDEX VALUE
// Return:操作成功返回 ok ，否则返回错误信息。
func LSet(rp *redis.Pool, key string, index int, value string) error {
	conn := rp.Get()
	defer conn.Close()

	_, err := conn.Do("LSET", key, index, value)
	return err
}

// Function:Lindex 命令用于通过索引获取列表中的元素
// Description:LINDEX KEY_NAME INDEX_POSITION
// Return:
func LIndex(rp *redis.Pool, key string, index int) ([]string, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("LINDEX", key, index)
	return redis.Strings(reply, err)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

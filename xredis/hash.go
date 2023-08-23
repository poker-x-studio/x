/*
功能：redis hash
时间：
说明：hash 是一个 string 类型的 field（字段） 和 value（值） 的映射表，hash 特别适合用于存储对象。
https://www.runoob.com/redis/redis-hashes.html
*/
package xredis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Field_value struct {
	Field string
	Value string
}

// HDel
// Function:Hdel 命令用于删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略。
// Description:HDEL KEY_NAME FIELD1.. FIELDN
// Return:被成功删除字段的数量，不包括被忽略的字段。
func HDel(rp *redis.Pool, key, field string) (int, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("HDEL", key, field)
	return redis.Int(reply, err)
}

// Function:Hexists 命令用于查看哈希表的指定字段是否存在。
// Description:HEXISTS KEY_NAME FIELD_NAME
// Return:如果哈希表含有给定字段，返回 1 。 如果哈希表不含有给定字段，或 key 不存在，返回 0 。
func HExists(rp *redis.Pool, key, field string) (int, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("HEXISTS", key, field)
	return redis.Int(reply, err)
}

// Function:Hget 命令用于返回哈希表中指定字段的值。
// Description:HGET KEY_NAME FIELD_NAME
// Return:返回给定字段的值。如果给定的字段或 key 不存在时，返回 nil 。
func HGet(rp *redis.Pool, key, field string) (string, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("HGET", key, field)
	return redis.String(reply, err)
}

// Function:Hgetall 命令用于返回哈希表中，所有的字段和值。
// Description:HGETALL KEY_NAME
// Return:以列表形式返回哈希表的字段及字段值。 若 key 不存在，返回空列表。
func HGetAll(rp *redis.Pool, key string) ([]*Field_value, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("HGETALL", key)
	if err != nil {
		Log().Error(err)
		return nil, err
	}
	var str_slice []string = make([]string, 0)
	for _, v := range reply.([]interface{}) {
		s := fmt.Sprintf("%s", v)
		str_slice = append(str_slice, s)
	}

	var rt []*Field_value = make([]*Field_value, 0)
	for i := 0; i < len(str_slice); i += 2 {
		item := &Field_value{
			Field: str_slice[i],
			Value: str_slice[i+1],
		}
		rt = append(rt, item)
	}
	return rt, nil
}

// Function:Hkeys 命令用于获取哈希表中的所有域（field）。
// Description:HKEYS key
// Return:包含哈希表中所有域（field）列表。 当 key 不存在时，返回一个空列表。
func HKeys(rp *redis.Pool, key string) ([]string, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("HKEYS", key)
	return redis.Strings(reply, err)
}

// Function:Hlen 命令用于获取哈希表中字段的数量。
// Description:HLEN KEY_NAME
// Return:哈希表中字段的数量。 当 key 不存在时，返回 0 。
func HLen(rp *redis.Pool, key string) (int, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("HLEN", key)
	return redis.Int(reply, err)
}

// Function:Hset 命令用于为哈希表中的字段赋值
// Description:HSET KEY_NAME FIELD VALUE
// Return:如果字段是哈希表中的一个新建字段，并且值设置成功，返回 1 。
//
//	如果哈希表中域字段已经存在且旧值已被新值覆盖，返回 0 。
func HSet(rp *redis.Pool, key, field, value string) (int, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("HSET", key, field, value)
	return redis.Int(reply, err)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

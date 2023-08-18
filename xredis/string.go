/*
功能：redis 字符串
说明：Redis 字符串数据类型的相关命令用于管理 redis 字符串值
*/
package xredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// 设置字符串
func SetString(rp *redis.Pool, key, value string) error {
	conn := rp.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}

// 设置字符串-有效期
func SetStringEx(rp *redis.Pool, key, value string, expire time.Duration) error {
	conn := rp.Get()
	defer conn.Close()

	expireSecond := int64(expire / time.Second)
	_, err := conn.Do("SETEX", key, expireSecond, value)
	return err
}

func GetString(rp *redis.Pool, key string) (string, error) {
	conn := rp.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

// 设置整数
func SetInt(rp *redis.Pool, key string, value int) error {
	conn := rp.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}

// 设置整数-有效期
func SetIntEx(rp *redis.Pool, key string, value int, expire time.Duration) error {
	conn := rp.Get()
	defer conn.Close()

	expireSecond := int64(expire / time.Second)
	_, err := conn.Do("SETEX", key, expireSecond, value)
	return err
}

func GetInt(rp *redis.Pool, key string) (int, error) {
	conn := rp.Get()
	defer conn.Close()

	return redis.Int(conn.Do("GET", key))
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

/*
功能：
说明：

参考文档 https://www.runoob.com/redis/redis-lists.html
*/
package xredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// OpenRedis 开启redis连接池
func OpenRedis(addr string, password string, db int) (*redis.Pool, error) {
	// 连接 redis
	rp := new(redis.Pool)
	rp.Dial = func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", addr)
		if err != nil {
			return nil, err
		}

		if password != "" {
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
		}

		if db != 0 {
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
		}

		return c, nil
	}

	rp.MaxActive = 20
	rp.MaxIdle = 20
	// redis服务默认超时时间300s
	rp.IdleTimeout = 30 * time.Second
	rp.Wait = true

	// TestNet
	conn := rp.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		rp.Close()
		return nil, err
	}

	return rp, nil
}

// CloseRedis 关闭redis连接池
func CloseRedis(rp *redis.Pool) {
	if rp != nil {
		rp.Close()
	}
}

// SetKeyExpire 设置key生命周期
func SetKeyExpire(rp *redis.Pool, key string, expire time.Duration) error {
	conn := rp.Get()
	defer conn.Close()

	expireSecond := int64(expire / time.Second)
	_, err := conn.Do("EXPIRE", key, expireSecond)
	return err
}

// DelKey 删除key
func DelKey(rp *redis.Pool, key string) error {
	conn := rp.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

// ExistsKey 是否存在key
func ExistsKey(rp *redis.Pool, key string) (int, error) {
	conn := rp.Get()
	defer conn.Close()

	reply, err := conn.Do("EXISTS", key)
	return redis.Int(reply, err)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

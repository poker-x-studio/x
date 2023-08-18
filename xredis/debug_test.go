/*
功能：测试单元
说明：
*/
package xredis

import (
	"fmt"
	"testing"
	"time"
)

var address = "127.0.0.1:6379"
var pwd = ""

func TestRedis(t *testing.T) {
	rds, err := OpenRedis(address, pwd, 0)
	if err != nil {
		t.Error(err)
		return
	}

	key := "testkey"
	value := 123456789
	err = SetIntEx(rds, key, value, time.Minute)
	if err != nil {
		t.Error(err)
		return
	}

	is_exists, err := ExistsKey(rds, "testkey")
	if is_exists > 0 {
		fmt.Println("存在")
		DelKey(rds, "testkey")
	}
	time.Sleep(time.Second * 5)
	fmt.Print(GetInt(rds, key))
}

func TestString(t *testing.T) {
	rds, err := OpenRedis(address, pwd, 0)
	if err != nil {
		t.Error(err)
		return
	}

	key := "key_string"
	value := 123456789
	err = SetIntEx(rds, key, value, time.Minute)
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 5)
	v, _ := GetInt(rds, key)
	fmt.Println(v)
}

func TestHash(t *testing.T) {
	rds, err := OpenRedis(address, pwd, 0)
	if err != nil {
		t.Error(err)
		return
	}

	key := "room_list"
	field1 := "room1"
	value1 := "value1"

	field2 := "room2"
	value2 := "value2"

	cnt, err := HSet(rds, key, field1, value1)
	if err != nil {
		return
	}
	t.Log(cnt)

	cnt, err = HSet(rds, key, field2, value2)
	if err != nil {
		return
	}
	t.Log(cnt)

	//is_exists, err := HExists(rds, key, field1)
	is_exists, err := HExists(rds, key, "xxxx")
	if err != nil {
		return
	}
	t.Log(is_exists)

	len, err := HLen(rds, key)
	if err != nil {
		return
	}
	t.Log(len)

	HGet(rds, key, field1)

	value_, err := HGet(rds, key, field1)
	if err != nil {
		return
	}
	t.Log(value_)

	//_, err = HDel(rds, key, field1)
	_, err = HGetAll(rds, key)

	field_list, err := HKeys(rds, key)
	if err != nil {
		return
	}
	t.Log(field_list)
}

// 测试非正常情况
func TestHashEx(t *testing.T) {
	rds, err := OpenRedis(address, pwd, 0)
	if err != nil {
		t.Error(err)
		return
	}

	key := "room_list1"
	field1 := "room1"
	//value1 := "value1"

	//field2 := "room2"
	//value2 := "value2"

	HGet(rds, key, field1)

	HGetAll(rds, key)

	HKeys(rds, key)

}

func TestList(t *testing.T) {
	rds, err := OpenRedis(address, pwd, 0)
	if err != nil {
		t.Error(err)
		return
	}

	key := "redis_list"
	value := "123456789"
	len, err := LPush(rds, key, value)

	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d", i)
		len, err := LPush(rds, key, str)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(len)
	}

	//var list []string
	str_list, err := LRange(rds, key, 0, int(len)-1)
	for k, v := range str_list {
		fmt.Println(k, v)
	}

	LIndex(rds, key, 100)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

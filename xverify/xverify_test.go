/*
功能：测试单元
说明：
*/
package xverify

import (
	"fmt"
	"testing"
)

func Test_account(t *testing.T) {
	const min_len = 4
	const max_len = 16
	account := "_dafadf123456789123456"
	matched, err := Verify_account(account, min_len, max_len)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)

	account = "_daf123456"
	matched, err = Verify_account(account, min_len, max_len)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)
}

func Test_pwd(t *testing.T) {
	const min_len = 4
	const max_len = 16
	pwd := "_dafadf123456789123456"
	matched, err := Verify_pwd(pwd, min_len, max_len)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)

	pwd = "_daf123456"
	matched, err = Verify_pwd(pwd, min_len, max_len)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

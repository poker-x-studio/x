/*
功能：测试单元
说明：
*/
package xcrypto

import (
	"fmt"
	"testing"
)

func Test_md5(t *testing.T) {
	fmt.Println(MD5([]byte("ddddd")))
}

func Test_aes(t *testing.T) {
	test_txt := "adafdafasf"

	encrypted_txt, err := AesEncrypt([]byte(test_txt))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(encrypted_txt))

	txt, err := AesDecrypt(encrypted_txt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(txt))

}

//-----------------------------------------------
//					the end
//-----------------------------------------------

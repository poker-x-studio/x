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

//-----------------------------------------------
//					the end
//-----------------------------------------------

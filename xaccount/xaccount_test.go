/*
功能：测试单元
说明：
*/
package xaccount

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	fmt.Println(Check_account("_dafadf123456789123456"))
	fmt.Println(Check_account("da_fadf"))
}
func Test_2(t *testing.T) {
	fmt.Println(Salt(5))
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

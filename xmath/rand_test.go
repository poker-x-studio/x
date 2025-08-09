/*
功能：测试单元
说明：
*/
package xmath

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	r, err := Intn(100)
	fmt.Println(r, err)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

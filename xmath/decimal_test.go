/*
功能：测试单元
说明：
*/
package xmath

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

// float64
func Test(t *testing.T) {
	fmt.Println(decimal.DivisionPrecision)
	decimal.DivisionPrecision = 6 //设置精度

	d1 := DECIMAL_MAKE_FLOAT32(2.3456)
	fmt.Println(d1.Neg().StringFixed(5))
	fmt.Println(d1.Neg().String())

}

//-----------------------------------------------
//					the end
//-----------------------------------------------

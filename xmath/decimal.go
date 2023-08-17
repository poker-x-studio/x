/*
功能：浮点数精度
说明：
*/
package xmath

import (
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

const (
	DECIMAL = 2 //浮点数精度
)

//decimal.DivisionPrecision = 3 //设置精度

// float64
func DECIMAL_2_FLOAT(d decimal.Decimal) float64 {
	f64, _ := d.Float64()
	return f64
}

func DECIMAL_MAKE_FLOAT32(f float32) decimal.Decimal {
	return decimal.NewFromFloat32(f)
}

func DECIMAL_MAKE(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f)
}

func DECIMAL_MAKE_INT32(n int32) decimal.Decimal {
	return decimal.NewFromInt32(n)
}

func DECIMAL_MAKE_INT(n int64) decimal.Decimal {
	return decimal.NewFromInt(n)
}

func DECIMAL_MAKE_STRING(s string) (decimal.Decimal, error) {
	return decimal.NewFromString(s)
}

func DECIMAL_MAKE_ZERO() decimal.Decimal {
	return decimal.NewFromFloat32(0.0)
}

//运算

// 加
func DECIMAL_ADD(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Add(d2)
}

func DECIMAL_ADD_EX(d1 decimal.Decimal, d2 decimal.Decimal) float64 {
	d3 := d1.Add(d2)
	return DECIMAL_2_FLOAT(d3)
}

// 减
func DECIMAL_SUB(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Sub(d2)
}

func DECIMAL_SUB_EX(d1 decimal.Decimal, d2 decimal.Decimal) float64 {
	d3 := d1.Sub(d2)
	return DECIMAL_2_FLOAT(d3)
}

// 乘
func DECIMAL_MUL(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Mul(d2)
}

// 除
func DECIMAL_DIV(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Div(d2)
}

// 浮点数转字符串
func Float_2_string(num float64) string {
	const decimal = DECIMAL //默认精度

	// 默认乘1
	d := float64(1.0)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去,小数点后无效的0也就不存在了
	f := math.Trunc(num*d) / d
	//两位小数
	return strconv.FormatFloat(f, 'f', decimal, 64)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

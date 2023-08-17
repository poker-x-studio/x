/*
功能：最小最大函数
说明：
*/
package xmath

func Min[T byte | int | int32 | int64 | float32 | float64](a T, b T) T {
	if a > b {
		return b
	}
	return a
}

func Max[T byte | int | int32 | int64 | float32 | float64](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

/*
功能：
说明：
*/
package xutils

// 是否包含
func Is_contains[T byte | int | int32 | int64](all []T, element T) bool {
	if all == nil {
		return false
	}
	for _, v := range all {
		if v == element {
			return true
		}
	}
	return false
}

// 插入
func Slice_insert[T byte | int](all []T, index int, data T) []T {
	new_all := make([]T, 0)
	if index == len(all) { //插入最后
		new_all = append(new_all, all...)
		new_all = append(new_all, data)
		return new_all
	}

	//插入开头/中间
	for i := 0; i < len(all); i++ {
		if i == index {
			new_all = append(new_all, data)
		}
		new_all = append(new_all, all[i])
	}
	return new_all
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

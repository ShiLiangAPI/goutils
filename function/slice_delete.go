package function

// SliceDeleteSelfValue list为基础数据，val为删除值
// 效率更高，但修改了自己切片的值
func SliceDeleteSelfValue[T comparable](list []T, val T) ([]T, bool) {
	j := 0
	isExist := false
	for _, v := range list {
		if v != val {
			list[j] = v
			j++
		} else {
			isExist = true
		}
	}

	return list[:j], isExist
}

// SliceDeleteCopyValue list为基础数据，val为删除值
// 效率没用上面的高，但不修改原切片
func SliceDeleteCopyValue[T comparable](list []T, val T) ([]T, bool) {
	temp := make([]T, 0, len(list))
	isExist := false
	for _, v := range list {
		if v != val {
			temp = append(temp, val)
		} else {
			isExist = true
		}
	}

	return temp, isExist
}

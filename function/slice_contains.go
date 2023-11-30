package function

// SliceContainsValue list为基础数据(全)，val为校验值
func SliceContainsValue[T comparable](list []T, val T) bool {

	for _, v := range list {
		if v == val {
			return true
		}
	}

	return false
}

// SliceContainsSlice list为基础数据(全)，valList为校验数组(部分)
func SliceContainsSlice[T comparable](list []T, valList []T) bool {

	temp := map[T]struct{}{}

	for _, v := range list {
		temp[v] = struct{}{}
	}

	for _, tv := range valList {
		_, ok := temp[tv]
		if !ok {
			return false
		}
	}

	return true
}

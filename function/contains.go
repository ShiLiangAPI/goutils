package function

func ValueInSlice[T comparable](list []T, val T) bool {
	temp := map[T]struct{}{}

	for _, v := range list {
		temp[v] = struct{}{}
	}

	if _, ok := temp[val]; !ok {
		return false
	}
	return true
}

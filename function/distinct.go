package function

// Distinct 去重
func Distinct[T comparable](el []T) []T {
	var result []T
	temp := map[T]struct{}{}

	for _, val := range el {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
			result = append(result, val)
		}
	}

	return result
}

// Different 差集 => 留下 el1 中不同的值
func Different[T comparable](el1 []T, el2 []T) []T {
	var result []T
	temp := map[T]struct{}{}

	for _, val := range el2 {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range el1 {
		if _, ok := temp[val]; !ok {
			result = append(result, val)
		}
	}

	return result
}

// Intersection 交集
func Intersection[T comparable](el1 []T, el2 []T) []T {
	var result []T
	temp := map[T]struct{}{}

	for _, val := range el2 {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range el1 {
		if _, ok := temp[val]; ok {
			result = append(result, val)
		}
	}

	return result
}

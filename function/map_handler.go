package function

// MapMerge map合并，同样的key值，以返回值为准
func MapMerge[T any](mapResp map[string]T, mapVal map[string]T) map[string]T {
	for key, value := range mapResp {
		mapVal[key] = value
	}

	return mapVal
}

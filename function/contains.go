package function

// ValueInSlice list为基础数据(全)，val为校验值
func ValueInSlice[T comparable](list []T, val T) bool {

	for _, v := range list {
		if v == val {
			return true
		}
	}

	return false
}

// SliceInSlice list为基础数据(全)，valList为校验数组(部分)
func SliceInSlice[T comparable](list []T, valList []T) (resMap map[T]bool) {

	temp := map[T]struct{}{}

	for _, v := range list {
		temp[v] = struct{}{}
	}

	for _, tv := range valList {
		_, ok := temp[tv]
		resMap[tv] = ok
	}

	return resMap
}

type SliceContains[T comparable] struct {
	SliceToMap map[T]struct{}
}

func NewSliceContains[T comparable](list []T) *SliceContains[T] {
	temp := map[T]struct{}{}

	for _, v := range list {
		temp[v] = struct{}{}
	}

	return &SliceContains[T]{
		SliceToMap: temp,
	}
}

func (obj *SliceContains[T]) Contains(val T) bool {
	_, ok := obj.SliceToMap[val]

	return ok
}

//func main() {
//	test := []int64{1, 2, 3, 5, 6}
//	c := NewSliceContains[int64](test)
//	fmt.Println(c.Contains(3))
//	fmt.Println(c.Contains(4))
//}

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
func SliceInSlice[T comparable](list []T, valList []T) bool {

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

// ContainsSlice 包含切片中的任意值
func (obj *SliceContains[T]) ContainsSlice(valList []T) bool {
	for _, val := range valList {
		_, ok := obj.SliceToMap[val]
		if ok {
			return true
		}
	}

	return false
}

//func main() {
//	test := []int64{1, 2, 3, 5, 6}
//	c := NewSliceContains[int64](test)
//	fmt.Println(c.Contains(3))
//	fmt.Println(c.Contains(4))
//}

package function

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

// Add 添加
func (obj *SliceContains[T]) Add(val T) {
	obj.SliceToMap[val] = struct{}{}

	return
}

// AddSlice 添加切片
func (obj *SliceContains[T]) AddSlice(valList []T) {
	for _, val := range valList {
		obj.Add(val)
	}

	return
}

//func main() {
//	test := []int64{1, 2, 3, 5, 6}
//	c := NewSliceContains[int64](test)
//	fmt.Println(c.Contains(3))
//	fmt.Println(c.Contains(4))
//}

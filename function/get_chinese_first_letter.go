package function

import "github.com/mozillazg/go-pinyin"

var pinyinRead *pinyin.Args

func InitPinyin() {
	read := pinyin.NewArgs()
	read.Style = pinyin.FirstLetter
	pinyinRead = &read
}

func GetChineseFirstLetter(chinese string) string {
	if pinyinRead == nil {
		InitPinyin()
	}
	zh := string([]rune(chinese)[0])

	return pinyin.Pinyin(zh, *pinyinRead)[0][0]
}

//func main() {
//	print(GetChineseFirstLetter("萌"))
//}

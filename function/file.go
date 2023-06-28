package function

import (
	"os"
	"path/filepath"
	"strings"
)

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {

		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// CheckFileExt 检查文件格式
func CheckFileExt(fileName string) bool {
	ext := filepath.Ext(fileName)
	for _, allowExt := range []string{"doc", "docx", "xls", "xlsx", "jpg", "png"} {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

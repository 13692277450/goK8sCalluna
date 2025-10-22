package util

import (
	"os"
	"strings"
)

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 20)
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	Path := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, folder := range dir {
		if folder.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(folder.Name()), suffix) { //匹配文件
			files = append(files, dirPth+Path+folder.Name())
		}
	}
	return files, nil
}

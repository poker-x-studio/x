/*
功能：路径
说明：
path = dir + filename
*/
package xpath

import (
	"fmt"
	"os"
	"path/filepath"
)

// 可执行文件-path
func Executable_path() string {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return path
}

// 可执行文件-dir
func Executable_dir() string {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return filepath.Dir(path)
}

// 可执行文件-文件名
func Executable_filename() string {
	pth, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return filepath.Base(pth)
}

// 判断所给路径文件/文件夹是否存在
func Exist(name string) bool {
	_, err := os.Stat(name) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 判断所给路径是否为文件夹
func Is_dir(name string) bool {
	fileinfo, err := os.Stat(name) //os.Stat获取文件信息
	if err != nil {
		return false
	}
	return fileinfo.IsDir()
}

// 判断所给路径是否为文件
func Is_file(name string) bool {
	return !Is_dir(name)
}

// 创建目录
func Mkdir(name string) error {
	if Exist(name) {
		return nil
	}
	//创建
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

/*
功能：路径
说明：
path = dir + file
*/
package xpath

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	path := Executable_path()
	fmt.Println(path)

	dir := Executable_dir()
	fmt.Println(dir)

	filename := Executable_filename()
	fmt.Println(filename)

	is_exist := Exist(path)
	fmt.Println("path:", path, "is_exist:", is_exist)

	is_exist = Exist(dir)
	fmt.Println("dir:", dir, "is_exist:", is_exist)

	path = "C:\\_test_proj\\go\\Telegram\\x1"
	is_exist = Exist(path)
	fmt.Println("path:", path, "is_exist:", is_exist)
}
func Test2(t *testing.T) {
	files := "E:\\data\\test.txt"
	paths, fileName := filepath.Split(files)
	fmt.Println(paths, fileName)      //获取路径中的目录及文件名 E:\data\  test.txt
	fmt.Println(filepath.Base(files)) //获取路径中的文件名test.txt
	fmt.Println(path.Ext(files))      //获取路径中的文件的后缀 .txt
}

func Test_filename(t *testing.T) {

	for i := 0; i < 10; i++ {
		filename := Rand_filename(10)
		fmt.Println(filename)
	}

	for i := 0; i < 10; i++ {
		filename := Date_filename()
		fmt.Println(filename)
	}

}

//-----------------------------------------------
//					the end
//-----------------------------------------------

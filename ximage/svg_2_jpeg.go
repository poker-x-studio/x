/*
功能：svg转jpeg
说明：
*/
package ximage

import (
	"fmt"
	"os"
)

// svg转jpeg
// 删除过渡的临时文件
func Svg_2_jpeg(content string) (jpg_filepath string, svg_filepath string, err error) {
	png_filepath := ""
	png_filepath, svg_filepath, err = Svg_2_png(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(png_filepath)

	jpg_filepath, err = Png_2_jpeg(png_filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//删除过渡的临时文件
	os.Remove(png_filepath)
	//fmt.Println(jpg_filepath)
	return
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

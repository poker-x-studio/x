/*
功能：svg存储
说明：
*/
package ximage

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"x/xpath"
)

// 存储svg文件
func svg_file_saver(content string) (string, error) {
	current_dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	//创建svg目录
	svg_dir := path.Join(current_dir, FOLDER_IMAGES_TMP)
	xpath.Mkdir(svg_dir)

	//日期文件名
	svg_filename := xpath.Date_filename() + EXT_SVG

	svg_filepath := path.Join(current_dir, FOLDER_IMAGES_TMP, svg_filename)
	file, err := os.OpenFile(svg_filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return "", err
	}
	defer file.Close()

	//写入文件
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return "", err
	}
	writer.Flush()
	return svg_filepath, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

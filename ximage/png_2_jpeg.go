/*
功能：png转jpeg
说明：
*/
package ximage

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

const (
	JPEG_QUALITY = 95 //jpeg品质
)

// Png_2_jpeg png转jpeg
func Png_2_jpeg(png_filepath string) (string, error) {
	file, err := os.Open(png_filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(png_filepath)
	out_filepath := strings.Replace(png_filepath, ext, EXT_JPEG, 1)

	// 输出文件
	out, err := os.Create(out_filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 将图像转换为JPEG格式并写入文件
	err = jpeg.Encode(out, img, &jpeg.Options{Quality: JPEG_QUALITY})
	if err != nil {
		return "", err
	}
	return out_filepath, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

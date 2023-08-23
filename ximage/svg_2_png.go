/*
功能：svg转png
说明：
*/
package ximage

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// Svg_2_png svg转png
func Svg_2_png(content string) (png_filepath string, svg_filepath string, err error) {
	svg_filepath, err = svg_file_saver(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	inkscape := get_inkscape()
	png_filepath, err = svg_2_png(inkscape, svg_filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(png_filepath)
	return
}

// svg_2_png 转png
func svg_2_png(inkscape_executable string, svg_filepath string) (string, error) {
	cmd := exec.Command(inkscape_executable, "--export-type=png", svg_filepath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	ext := filepath.Ext(svg_filepath)
	png_filepath := strings.Replace(svg_filepath, ext, EXT_PNG, 1)

	return png_filepath, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

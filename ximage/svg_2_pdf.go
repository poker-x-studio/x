/*
功能：svg转pdf
说明：
*/
package ximage

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// Svg_2_pdf svg转pdf
func Svg_2_pdf(content string) (pdf_filepath string, svg_filepath string, err error) {
	svg_filepath, err = svg_file_saver(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	inkscape := get_inkscape()
	pdf_filepath, err = svg_2_pdf(inkscape, svg_filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(png_filepath)
	return
}

// svg_2_pdf 转pdf
func svg_2_pdf(inkscape_executable string, svg_filepath string) (string, error) {
	cmd := exec.Command(inkscape_executable, "--export-type=pdf", svg_filepath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	ext := filepath.Ext(svg_filepath)
	pdf_filepath := strings.Replace(svg_filepath, ext, EXT_PDF, 1)

	return pdf_filepath, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

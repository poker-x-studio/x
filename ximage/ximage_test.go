/*
功能：测试单元
说明：
*/
package ximage

import (
	"fmt"
	"os/exec"
	"testing"
)

func Test_x(t *testing.T) {
	content := content()
	svg_filepath, err := svg_file_saver(content)
	if err != nil {
		fmt.Println(err)
	}

	inkscape := get_inkscape()
	png_filepath, err := svg_2_png(inkscape, svg_filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(png_filepath)

	pdf_filepath, err := svg_2_pdf(inkscape, svg_filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pdf_filepath)

	jpg_filepath, err := Png_2_jpeg(png_filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jpg_filepath)
}

func Test_convert_to_png(t *testing.T) {
	inkscape := get_inkscape()
	cmd := exec.Command(inkscape, "--export-type=jpg", "test.svg")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Test_convert_to_pdf(t *testing.T) {
	inkscape := get_inkscape()
	cmd := exec.Command(inkscape, "--export-type=pdf", "test.svg")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func content() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<svg width="594px" height="975px" viewBox="0 0 396 650" version="1.1"
	xmlns="http://www.w3.org/2000/svg"
	xmlns:xlink="http://www.w3.org/1999/xlink">
	<defs>
		<radialGradient id="c">
			<stop offset="0%" stop-color="#FFF"/>
			<stop offset="20%" stop-color="#FFF"/>
			<stop offset="100%" stop-color="#007ec5"/>
		</radialGradient>
		<g id="ball">
			<circle cx="85" cy="15" r="13.5" fill="url(#c)"/>
			<circle cx="130" cy="15" r="13.5" fill="url(#c)"/>
			<circle cx="175" cy="15" r="13.5" fill="url(#c)"/>
			<text x="102" y="21" font-size="18" fill="#000">+</text>
			<text x="147" y="21" font-size="18" fill="#000">+</text>
			<text x="194" y="21" font-size="18" fill="#000">=</text>
		</g>
		<g id="rbo">
			<rect x="0" y="0" width="396" height="30" fill="#F5FAFF"/>
			<use xlink:href="#ball"/>
		</g>
		<g id="rbe">
			<rect x="0" y="0" width="396" height="30" fill="#FFF"/>
			<use xlink:href="#ball"/>
		</g>
	</defs>
	<g transform="translate(0,0)">
		<rect x="0" y="0" width="396" height="50" fill="#6664C5"/>
		<text x="12" y="32" font-size="18" fill="#FFF">回合</text>
		<text x="94" y="32" font-size="18" fill="#FFF">结果</text>
		<text x="200" y="32" font-size="18" fill="#FFF">特码</text>
		<text x="250" y="32" font-size="18" fill="#FFF">双面</text>
		<text x="300" y="32" font-size="18" fill="#FFF">极值</text>
		<text x="350" y="32" font-size="18" fill="#FFF">形态</text>
	</g>
	<g transform="translate(0,0)">
		<g transform="translate(0,50)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,80)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">2345678</text>
			<text x="81" y="21" font-size="16" fill="#000">4</text>
			<text x="126" y="21" font-size="16" fill="#000">5</text>
			<text x="171" y="21" font-size="16" fill="#000">6</text>
			<text x="210" y="21" font-size="16" fill="#000">16</text>
			<text x="250" y="21" font-size="16" fill="#F00">大</text>
			<text x="270" y="21" font-size="16" fill="#5468a9">单</text>
			<text x="300" y="21" font-size="16" fill="#5468a9">极大</text>
			<text x="350" y="21" font-size="16" fill="#00ca3d">杂六</text>
		</g>
		<g transform="translate(0,110)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">3456789</text>
			<text x="81" y="21" font-size="16" fill="#000">7</text>
			<text x="126" y="21" font-size="16" fill="#000">8</text>
			<text x="171" y="21" font-size="16" fill="#000">9</text>
			<text x="210" y="21" font-size="16" fill="#000">26</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#5468a9">单</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,140)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">4567890</text>
			<text x="81" y="21" font-size="16" fill="#000">9</text>
			<text x="126" y="21" font-size="16" fill="#000">9</text>
			<text x="171" y="21" font-size="16" fill="#000">9</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#F00">大</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,170)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">1</text>
			<text x="171" y="21" font-size="16" fill="#000">1</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#5468a9">单</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,200)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">2</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">2</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#F00">大</text>
			<text x="270" y="21" font-size="16" fill="#5468a9">单</text>
			<text x="300" y="21" font-size="16" fill="#c00014">极小</text>
			<text x="350" y="21" font-size="16" fill="#00ca3d">杂六</text>
		</g>
		<g transform="translate(0,230)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">3</text>
			<text x="126" y="21" font-size="16" fill="#000">3</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#F00">大</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,260)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#5468a9">极大</text>
			<text x="350" y="21" font-size="16" fill="#00ca3d">杂六</text>
		</g>
		<g transform="translate(0,290)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#5468a9">单</text>
			<text x="300" y="21" font-size="16" fill="#5468a9">极大</text>
			<text x="350" y="21" font-size="16" fill="#00ca3d">杂六</text>
		</g>
		<g transform="translate(0,320)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,350)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,380)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#F00">大</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#f3a249">对子</text>
		</g>
		<g transform="translate(0,410)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#5468a9">单</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,440)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#c00014">极小</text>
			<text x="350" y="21" font-size="16" fill="#00ca3d">杂六</text>
		</g>
		<g transform="translate(0,470)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#F00">大</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,500)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,530)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
		<g transform="translate(0,560)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#F00">大</text>
			<text x="270" y="21" font-size="16" fill="#5468a9">单</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#F0F">豹子</text>
		</g>
		<g transform="translate(0,590)">
			<use xlink:href="#rbe"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#F0F">豹子</text>
		</g>
		<g transform="translate(0,620)">
			<use xlink:href="#rbo"/>
			<text x="5" y="21" font-size="16" fill="#000">1234567</text>
			<text x="81" y="21" font-size="16" fill="#000">1</text>
			<text x="126" y="21" font-size="16" fill="#000">2</text>
			<text x="171" y="21" font-size="16" fill="#000">3</text>
			<text x="210" y="21" font-size="16" fill="#000">6</text>
			<text x="250" y="21" font-size="16" fill="#5468a9">小</text>
			<text x="270" y="21" font-size="16" fill="#F00">双</text>
			<text x="300" y="21" font-size="16" fill="#888">无</text>
			<text x="350" y="21" font-size="16" fill="#6300ad">顺子</text>
		</g>
	</g>
</svg>
`
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

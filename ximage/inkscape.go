/*
功能：
说明：
1 https://inkscape.org/ 下载 inkscape，跨平台

2 将输入 SVG（in.svg）导出为 PDF（out.pdf）格式：

	inkscape --export-filename=out.pdf in.svg

3 将输入文件（in1.svg, in2.svg）导出为 PNG 格式，保留原始名称（in1.png, in2.png）：

	inkscape --export-type=png in1.svg in2.svg
*/
package ximage

import (
	"runtime"

	"github.com/poker-x-studio/x/xutils"
)

const (
	Inkscape_windows = "C:/Program Files/Inkscape/bin/inkscape"
	Inkscape_ubuntu  = "inkscape"
)

// 得到路径
func get_inkscape() string {
	inkscape := Inkscape_windows
	if runtime.GOOS == xutils.WINDOWS {
		inkscape = Inkscape_windows
	} else if runtime.GOOS == xutils.LINUX {
		inkscape = Inkscape_ubuntu
	}
	return inkscape
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

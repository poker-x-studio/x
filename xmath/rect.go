/*
功能：几何图形-区域定义
说明：
*/
package xmath

import "fmt"

// Rect 区域
type Rect struct {
	X_left   int
	X_right  int
	Y_top    int
	Y_bottom int
}

// NewRect 构造
func NewRect(left int, top int, width int, height int) *Rect {
	rct := &Rect{
		X_left:   left,
		X_right:  left + width,
		Y_top:    top,
		Y_bottom: top + height,
	}
	return rct
}

// NewRectWithCopy 构造
func NewRectWithCopy(r *Rect) *Rect {
	if r == nil {
		return nil
	}
	return NewRect(r.X_left, r.Y_top, r.Witdh(), r.Height())
}

// Center_X x 中心
func (r Rect) Center_X() int {
	return (r.X_left + r.X_right) / 2
}

// Center_y y 中心
func (r Rect) Center_y() int {
	return (r.Y_top + r.Y_bottom) / 2
}

// Witdh 高度
func (r Rect) Witdh() int {
	return (r.X_right - r.X_left)
}

// Height 高度
func (r Rect) Height() int {
	return (r.Y_bottom - r.Y_top)
}

// X_move x方向移动
func (r *Rect) X_move(x_space int) *Rect {
	r.X_left += x_space
	r.X_right += x_space
	return r
}

// X_move_to x方向移动到
func (r *Rect) X_move_to(new_x int) *Rect {
	w := r.Witdh()
	r.X_left = new_x
	r.X_right = r.X_left + w
	return r
}

// Update_width 改变宽度[左上不变]
func (r *Rect) Update_width(new_width int) *Rect {
	r.X_right = r.X_left + new_width
	return r
}

// Y_move y方向移动
func (r *Rect) Y_move(y_space int) *Rect {
	r.Y_top += y_space
	r.Y_bottom += y_space
	return r
}

// Reposition 再定位(宽度和高度不变的情况)
func (r *Rect) Reposition(x_left int, y_top int) *Rect {
	w := r.Witdh()
	h := r.Height()

	r.X_left = x_left
	r.Y_top = y_top
	r.X_right = r.X_left + w
	r.Y_bottom = r.Y_top + h
	return r
}

// String 文字描述
func (r *Rect) String() string {
	return fmt.Sprintf("[left:%d,top:%d,right:%d,bottom:%d]", r.X_left, r.Y_top, r.X_right, r.Y_bottom)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

/*
功能：生命周期
说明：
1 输出函数运行时间[单位秒]
*/
package xdebug

import (
	"testing"
)

func Test_lifetime(t *testing.T) {
	Init_logrus()

	var lifetime LifeTime
	lifetime.Start()
	defer lifetime.End()
}

func Test_goid(t *testing.T) {

}

//-----------------------------------------------
//					the end
//-----------------------------------------------

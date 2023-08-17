/*
功能：测试单元
说明：
*/
package xrpc

import "testing"

func Test(t *testing.T) {
	NewServer("service1", "localhost:8021")
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

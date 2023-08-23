/*
功能：base64
说明：对称加密算法
*/
package xcrypto

import "encoding/base64"

// Base64_encode 编码
func Base64_encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64_decode 解码
func Base64_decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

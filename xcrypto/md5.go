/*
功能：md5-32位
说明：非对称加密算法
*/
package xcrypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 md5加密,32位输出
func MD5(datas []byte) string {
	h := md5.New()
	h.Write(datas)
	return hex.EncodeToString(h.Sum(nil))
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

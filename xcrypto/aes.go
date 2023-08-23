/*
功能：aes - Advanced Encryption Standard (AES)
说明：对称加密算法
AES-128
*/
package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

const (
	AES_KEY = "27ad84xd62e27fe6" //128bit[16byte]
)

// AesEncrypt AES加密
func AesEncrypt(origData []byte) ([]byte, error) {
	key := aes_key()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt AES解密
func AesDecrypt(crypted []byte) ([]byte, error) {
	key := aes_key()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	return origData, nil
}

func aes_key() []byte {
	return []byte(AES_KEY)
}

// @brief:填充明文
func pKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// @brief:去除填充数据
func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//-----------------------------------------------
//					the end
//-----------------------------------------------

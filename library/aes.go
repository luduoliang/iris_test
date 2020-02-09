package library

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

//Aes CBC 加密
func AesCBCEncrypt(plainText []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	paddedText := PKCS5Padding(plainText, aes.BlockSize)

	encrypted := make([]byte, aes.BlockSize+len(paddedText))
	if len(iv) == 0 {
		iv = []byte(encrypted[:aes.BlockSize])
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return "", err
		}
	} else {
		encrypted = make([]byte, len(paddedText))
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	if len(encrypted) > len(paddedText) {
		blockMode.CryptBlocks(encrypted[aes.BlockSize:], paddedText)
	} else {
		blockMode.CryptBlocks(encrypted, paddedText)
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

//Aes CBC 解密
func AesCBCDecrypt(encodeStr []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	base64DecodeStr, err := base64.StdEncoding.DecodeString(string(encodeStr))
	if len(iv) == 0 {
		iv = base64DecodeStr[:aes.BlockSize]
		base64DecodeStr = base64DecodeStr[aes.BlockSize:]
	}

	if err != nil /*|| len(base64DecodeStr) <= aes.BlockSize || len(base64DecodeStr)%aes.BlockSize != 0*/ {
		return nil, errors.New("failed to decode base64")
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(base64DecodeStr))
	blockMode.CryptBlocks([]byte(plainText), []byte(base64DecodeStr))
	return PKCS5UnPadding(plainText), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
package crypro

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Base64Decode ..
func Base64Decode(data string) []byte {
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil
	}
	return res
}

// Base64Encode ..
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// EnctyprAES ..
func EnctyprAES(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	data = paddingPKCS5(data, blockSize)
	cipherText := make([]byte, blockSize+len(data))
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], data)
	return cipherText
}

// DecryptAES ...
func DecryptAES(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return nil
	}
	iv := data[:blockSize]
	data = data[blockSize:]
	if len(data)%blockSize != 0 {
		return nil
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return unpaddingPKS5(data)
}

func unpaddingPKS5(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unpadding := int(data[length-1])
	if length < unpadding {
		return nil
	}
	return data[:(length - unpadding)]
}

func paddingPKCS5(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	paddText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddText...)
}

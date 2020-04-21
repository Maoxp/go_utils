package maes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 加密 aes_128_cbc
func Encrypt(encryptStr string, key []byte, iv string) (string, error) {
	src := []byte(encryptStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	src = pkcs5Padding(src, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return base64.URLEncoding.EncodeToString(dst), nil
}

// 解密
func Decrypt(encryptedData string, key []byte, iv string) (string, error) {
	src, err := base64.URLEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	_iv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, _iv)
	dst := make([]byte, len(src))

	blockMode.CryptBlocks(dst , src)
	dst = pkcs5UnPadding(dst)
	return string(dst), nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs5UnPadding(decrypted []byte) []byte {
	length := len(decrypted)
	unPadding := int(decrypted[length-1])
	return decrypted[:(length - unPadding)]
}

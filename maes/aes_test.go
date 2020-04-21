package maes

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestEncrypt(t *testing.T) {
	var src = "hello, world"
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	var iv = "123456789aasdfhk"

	encryptedData, err := Encrypt(src, key, iv)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(encryptedData)

}

func TestDecrypt(t *testing.T) {
	var src = "-AbMAFrC0XNZiQZ4_mDQBg=="
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	var iv = "123456789aasdfhk"

	data, err := Decrypt(src, key, base64.StdEncoding.EncodeToString([]byte(iv)))
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(data)
}

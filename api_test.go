package shanyan

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestHmacSHA256(t *testing.T) {
	secret := "asdqwaswrasdqwer"
	data := "qwdqwdqwd"

	t.Log(HmacSHA256(secret, data))
}

func TestAesDecrypt(t *testing.T) {
	//key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
	var aeskey = []byte("12345678abcdefgh")
	pass := []byte("vdncloud123456")
	xpass, err := AesEncrypt(pass,aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n",pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		t.Error(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("解密后:%s\n", tpass)
}

func TestMapSort(t *testing.T) {
	data := map[string]string{
		"token": 		"asdasd123",
		"appId": 		"asdqwdsd",
		"test":         "asdasd",
	}

	t.Log(MapSort(data))
}

func TestClient_Sign(t *testing.T) {
	data := map[string]string{
		"token": 		"asdasd123",
		"appId": 		"asdqwdsd",
		"test":         "asdasd",
	}

	client := NewClient("appID", "appKEY")
	t.Log(client.Sign(MapSort(data)))
}

func TestClient_MobileQuery(t *testing.T) {
	client := NewClient("appID", "appKEY")
	mobile, err := client.MobileQuery("token")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(mobile)
}

//@brief:AES加密
func AesEncrypt(origData, key []byte) ([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData,blockSize)
	blockMode := cipher.NewCBCEncrypter(block,key[:blockSize])	//初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted,origData)
	return crypted, nil
}

//@brief:填充明文
func PKCS5Padding(plaintext []byte, blockSize int) []byte{
	padding := blockSize-len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(plaintext,padtext...)
}
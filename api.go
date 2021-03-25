package shanyan

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/levigross/grequests"
	"sort"
	"strconv"
)

const (
	HOST = "https://api.253.com"
	OPEN_MOBILE_QUERY = "/open/flashsdk/mobile-query"
)

type Client struct {
	appID  string
	appKey string
}

func NewClient(appID, appKey string) *Client {
	return &Client{appID: appID, appKey: appKey}
}

func (c Client) MobileQuery(token string) (string, error)  {
	req := map[string]string{
		"appId": 		c.appID,
		"token": 		token,
	}

	// 签名
	sign := c.Sign(MapSort(req))
	req["sign"] = sign

	options := &grequests.RequestOptions{
		Params: req,
	}

	resp, err := grequests.Post(HOST + OPEN_MOBILE_QUERY, options)
	if err != nil {
		return "", err
	}

	var mobileQueryResponse MobileQueryResponse
	err = resp.JSON(&mobileQueryResponse)
	if err != nil {
		return "", err
	}

	if mobileQueryResponse.BaseResponse.Code != strconv.Itoa(MSG_SUCCESS) {
		return "", errors.New(mobileQueryResponse.BaseResponse.Message)
	}

	// 解密出手机号
	mobile, err := AesDecrypt([]byte(mobileQueryResponse.Data.MobileName), []byte(c.appKey))
	if err != nil {
		return "", err
	}

	return string(mobile), nil
}

// shanyan签名
func (c Client) Sign(data string) string {
	return HmacSHA256(c.appKey, data)
}

// map排序
func MapSort(barVal map[string]string) string {
	data := ""
	keys := make([]string, len(barVal))

	// 排序
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	for _, k := range keys {
		data += k + barVal[k]
	}

	return data
}

// 加密
func HmacSHA256(secret, data string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}

//@brief:AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])	//初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)

	return origData, nil
}

//@brief:去除填充数据
func PKCS5UnPadding(origData []byte) []byte{
	length := len(origData)
	unpadding := int(origData[length-1])

	return origData[:(length - unpadding)]
}
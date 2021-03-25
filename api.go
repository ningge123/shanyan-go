package shanyan

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
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

	if mobileQueryResponse.Code != strconv.Itoa(MSG_SUCCESS) {
		return "", errors.New(mobileQueryResponse.BaseResponse.Message)
	}

	// 解密出手机号
	return AesDecrypt(mobileQueryResponse.Data.MobileName, c.appKey), nil
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
func AesDecrypt(crypted, key string) string {
	hash := md5.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	block, _ := aes.NewCipher([]byte(hashString[:16]))
	ecb := cipher.NewCBCDecrypter(block, []byte(hashString[16:]))
	source, _ := hex.DecodeString(crypted)
	decrypted := make([]byte, len(source))
	ecb.CryptBlocks(decrypted, source)

	return string(PKCS5UnPadding(decrypted))
}

//@brief:去除填充数据
func PKCS5UnPadding(origData []byte) []byte{
	length := len(origData)
	unpadding := int(origData[length-1])

	return origData[:(length - unpadding)]
}
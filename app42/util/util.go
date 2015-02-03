package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

type Param struct {
	ApiKey    string `json:"apiKey"`
	Version   string `json:"version"`
	TimeStamp string `json:"timeStamp"`
}

func Sign(secretKey, params string) string {
	sortedParams := sortConvert(params)
	signature := computeHMAC(secretKey, sortedParams)
	return signature
}

func sortConvert(params string) string {
	var sortedParams string

	var pro Param

	_ = json.Unmarshal([]byte(params), &pro)

	fmt.Println("pro==============>", reflect.TypeOf(pro))

	sortedParams = sortedParams + "apiKey" + pro.ApiKey + "timeStamp" + pro.TimeStamp + "version" + pro.Version

	return sortedParams
}

func computeHMAC(secretKey, sortedParams string) string {
	key := []byte(secretKey)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(sortedParams))

	siganture := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return url.QueryEscape(siganture)
}

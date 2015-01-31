package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strings"
	"time"
)

type Param struct {
	ApiKey    string `json:"apiKey"`
	Version   string `json:"version"`
	TimeStamp string `json:"timeStamp"`
}

func TimeStamp() string {
	currentTime := time.Now().UTC()

	formatedTime := currentTime.Format(time.RFC3339)

	dateTime := strings.Split(formatedTime, "T")

	stampMilli := currentTime.Format(time.StampMilli)

	timeArray := strings.Split(stampMilli, " ")

	s := []string{dateTime[0], timeArray[2]}
	final := strings.Join(s, "T")

	return final + "Z"
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

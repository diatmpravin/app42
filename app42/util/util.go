package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"net/url"
	"sort"
)

var p map[string]string

func Sign(secretKey, params string) string {
	sortedParams := sortConvert(params)
	signature := computeHMAC(secretKey, sortedParams)
	return signature
}

func sortConvert(params string) string {
	var sortedParams string

	p = make(map[string]string)

	_ = json.Unmarshal([]byte(params), &p)

	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		sortedParams = sortedParams + k + p[k]
	}

	return sortedParams
}

func computeHMAC(secretKey, sortedParams string) string {
	key := []byte(secretKey)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(sortedParams))

	siganture := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return url.QueryEscape(siganture)
}

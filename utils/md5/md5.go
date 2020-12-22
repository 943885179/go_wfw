package mzjmd5

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5 生成32位MD5
func MD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

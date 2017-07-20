package alivod

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
    "crypto/hmac"
    "bytes"
)


/**
 * 计算签名
 */
func sha1Sign(origData string, privateKey string) (string) {
    var buf bytes.Buffer
    defer buf.Reset()
    buf.WriteString("GET&%2F&")
    buf.WriteString(origData)
    fmt.Println(buf.String())
    //hmac ,use sha1
    key := []byte(privateKey+"&")
    mac := hmac.New(sha1.New, key)
    mac.Write(buf.Bytes())
    //fmt.Println(string(mac.Sum(nil)))
    data := base64.StdEncoding.EncodeToString(mac.Sum(nil))
    fmt.Println(string(data))
    return string(data)
}

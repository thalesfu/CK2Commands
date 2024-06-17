package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	jsonData := "{\"tplId\":\"106355\",\"isDepositedTemplate\":true,\"tplType\":1,\"tplVersion\":12345,\"recordId\":56791,\"replaceableJson\":[{\"main\":\"\"}],\"ossId\":\"8606b3a09a6976ea3dbae9c82468ee08\",\"notifyUrl\":\"http://www.baidu.com\",\"uploadKey\":\"vetest1.mp4\"}"
	productid := "4808cbd622"
	accessKey := "ea7ca35f6b785cd65dec30a94178f388"
	accessSecret := "b2109626b2c117e96d8c11e20d5137f0"

	data := jsonData + "&&" + productid + "&&" + accessKey + "&&" + getMD5Hash(accessSecret)

	fmt.Printf(data)
	fmt.Println()

	token := getMD5Hash(data)
	fmt.Println(token)
	fmt.Println(productid + ":" + accessKey + ":" + token)

}

// 计算字符串的MD5哈希值
func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

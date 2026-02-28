package ecerp

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

// GenerateSign 根据易仓ERP签名规则生成MD5签名
//
// 签名规则:
//  1. 收集除 sign 外的所有参数
//  2. 过滤空值参数
//  3. 按 key 字母序排序
//  4. 拼接为 key1=value1&key2=value2...
//  5. 末尾直接追加 app_secret
//  6. 计算 32 位小写 MD5
func GenerateSign(params map[string]string, appSecret string) string {
	// 过滤空值并收集key
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if k == "sign" || v == "" {
			continue
		}
		keys = append(keys, k)
	}

	// 按字母序排序
	sort.Strings(keys)

	// 拼接参数
	var buf strings.Builder
	for i, k := range keys {
		if i > 0 {
			buf.WriteString("&")
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(params[k])
	}

	// 追加 app_secret
	buf.WriteString(appSecret)

	// 计算MD5
	hash := md5.Sum([]byte(buf.String()))
	return fmt.Sprintf("%x", hash)
}

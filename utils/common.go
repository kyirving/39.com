package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
	"sort"
	"strings"
	"time"
)

// 生成唯一ID
func GenerateUniqueID(prefix string) string {
	t := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s", prefix, t)
}

func GetFnNameWithLine(skip int) (fn_name string, line int) {
	pc, _, line, ok := runtime.Caller(skip)
	if ok {
		fn_name = runtime.FuncForPC(pc).Name()
	}
	return
}

func Createsign(params map[string]interface{}, secretKey string) string {
	// 对参数进行排序
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if k == "sign" || k == "SIGN" || k == "Sign" || v == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var signStr strings.Builder
	for _, k := range keys {
		signStr.WriteString(fmt.Sprintf("%s=%v&", k, params[k]))
	}
	signStr.WriteString("key=" + secretKey)
	h := md5.New()
	h.Write([]byte(signStr.String()))
	sign := hex.EncodeToString(h.Sum(nil))
	return strings.ToLower(sign)
}

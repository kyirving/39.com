package utils

import (
	"fmt"
	"runtime"
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

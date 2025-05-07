package utils

import (
	"fmt"
	"time"
)

// 生成唯一ID
func GenerateUniqueID(prefix string) string {
	t := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s", prefix, t)
}

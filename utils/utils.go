package utils

import (
	"fmt"
	"strings"
)

// SetSQLFormat  设置SQL语句格式
func SetSQLFormat(fsql string, args ...interface{}) string {
	sql := fsql
	for i := 0; i < len(args); i++ {
		key := fmt.Sprintf("{%d}", i)

		val := fmt.Sprint(args[i])

		sql = strings.Replace(sql, key, val, -1)
	}

	return sql
}

// Error 错误信息
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

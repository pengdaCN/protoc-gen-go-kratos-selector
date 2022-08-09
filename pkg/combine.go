package pkg

import "strings"

// CombineOperation Deprecated: kratos新版更新了http的生成方式
func CombineOperation(op string) string {
	idx := strings.LastIndex(op, "/")
	if idx < 0 {
		return op
	}

	if len(op)-1 == idx {
		return op
	}

	bs := []byte(op)
	bs[idx+1] = strings.ToLower(op[idx+1 : idx+2])[0]

	return string(bs)
}

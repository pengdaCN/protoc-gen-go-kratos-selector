package pkg

import "strings"

func CombineOperation(op string) string {
	idx := strings.LastIndex(op, "/")
	if idx < 0 {
		return op
	}

	if len(op)-1 == idx {
		return op
	}

	bs := []byte(op)
	bs[idx+1] = strings.ToUpper(op[idx+1 : idx+2])[0]

	return string(bs)
}

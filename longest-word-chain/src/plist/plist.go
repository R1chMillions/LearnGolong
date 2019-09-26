package plist

import "container/list"

func GetListValue(l *list.List, n int) interface{} {
	res := l.Front()
	if n <= 0 && l.Len() < n {
		return res
	}
	for i := 0; i < n; i++ {
		res = res.Next()
	}
	return res.Value
}

func HasValue(l *list.List, val int) bool {
	for item := l.Front(); item != nil; item = item.Next() {
		if val == item.Value {
			return true
		}
	}
	return false
}

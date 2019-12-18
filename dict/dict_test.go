package main

import "testing"

func TestTrans(t *testing.T) {
	query1 := "good"
	query2 := "你好"
	query3 := "JSON 是一种轻量级的数据交换格式，常用作前后端数据交换，Go 在 encoding/json 包中提供了对 JSON 的支持。"
	Trans(query1)
	Trans(query2)
	Trans(query3)
}

func TestTruncate(t *testing.T) {
	querys := []struct{
		in string
		expected string
	}{
		{"你好", "你好"},
		{"hello", "hello"},
		{
			"JSON 是一种轻量级的数据交换格式，常用作前后端数据交换",
			"JSON 是一种轻量29常用作前后端数据交换",
		},
	}
	for _, tt := range querys {
		actual := Truncate(tt.in)
		if actual != tt.expected {
			t.Error(tt.in, " but: " ,tt.expected)
		}
	}
}

func BenchmarkTrans(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trans("hello")
	}
}

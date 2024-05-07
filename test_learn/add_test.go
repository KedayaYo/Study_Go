/**
 * @Author Kedaya
 * @Date 2024/5/5 16:38:00
 * @Desc
 **/
package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println("test add1")
	if testing.Short() {
		t.Skip("Short模式下跳过")
	}
	re := add(1, 2)
	if re != 3 {
		t.Errorf("expect %d, actual %d", 3, re)
		t.Fatal("err")
	}
	t.Log("success")

}

func TestAdd2(t *testing.T) {
	fmt.Println("test add2")
	if testing.Short() {
		t.Skip("Short模式下跳过") // 以时间短的方式运行测试  go test -short  耗时长的test 判断后面的不运行
	}

	dataset := []struct {
		a      int
		b      int
		expect int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{-9, 9, 0},
	}
	for _, v := range dataset {
		re := add(v.a, v.b)
		if re != v.expect {
			t.Errorf("expect %d, actual %d", v.expect, re)
			t.Fatal("err")
		}
	}
	t.Log("success")
}

// go test -bench=".*"
const num = 10000

func BenchmarkStringf(b *testing.B) {
	fmt.Println("benchmark stringf")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < num; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.ResetTimer()
}

func BenchmarkStringConv(b *testing.B) {
	fmt.Println("benchmark stringConv")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < num; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	b.ResetTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	fmt.Println("benchmark stringAdd")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < num; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.ResetTimer()
}

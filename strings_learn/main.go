package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

/**
 * @Title lengthOfNonRepeatingSubStr
 * @Description 无重复字符的最长子串
 * @Author Kedaya
 * @Time 2024-05-03 23:10:55
 * @Param s
 * @Return int
 **/
func lengthOfNonRepeatingSubStr(s string) int {
	/*
		lastOccurred := make(map[rune]int)：
		这一行初始化了一个哈希表 lastOccurred，用来存储每个字符最近一次出现的索引。键是字符，值是该字符最近一次出现的索引位置。
		start := 0：
		这定义了当前考虑的子串的起始索引，起始时为 0。
		maxLength := 0：
		这用于存储遍历字符串过程中发现的最长无重复字符子串的长度。
		for i, ch := range s：
		这是一个循环，遍历字符串 s 中的每个字符及其索引。i 是字符的索引，ch 是字符本身。
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start：
		这个条件检查当前字符 ch 是否之前出现过，并且出现的位置 lastI 是否在当前考虑的子串的起始位置 start 之后（或等于 start）。如果是这样，说明我们需要调整子串的起始位置 start 到 lastI + 1，以确保子串中没有重复的字符。
		if i-start+1 > maxLength：
		这个条件用来更新找到的最长无重复子串的长度。i-start+1 是当前考虑的子串的长度（因为子串是从 start 到 i）。如果这个长度大于之前记录的 maxLength，则更新 maxLength。
		lastOccurred[ch] = i：
		在哈希表中更新当前字符 ch 的最新索引位置 i。
		return maxLength：
		在循环结束后，返回找到的最长无重复子串的长度。
	*/
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range s {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func lastIndexRune(s *string, str string) int {
	if s == nil || str == "" {
		return -1 // 如果输入无效或空，返回-1
	}

	a := []rune(*s)
	targetRune, size := utf8.DecodeRuneInString(str)
	if size == 0 { // 检查解码是否成功
		return -1
	}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == targetRune {
			return i // 返回最后一个匹配字符的索引
		}
	}

	return -1 // 如果没有找到，返回-1
}

func main() {
	fmt.Println("查找无重复字符的最长子串的长度")
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))

	// 转义字符
	fmt.Println("转义字符")
	courseName := "Go\"语言\""
	fmt.Println(courseName)
	courseName = `Go"语言"`
	fmt.Println(courseName)

	/*
		格式化输出
		%v: 按照类型转化
		%#v: 类型+值
		%T: 类型
	*/
	fmt.Println("格式化输出")
	username := "Entic_Kedaya_浩浩"
	age := 20
	address := "上海"
	mobile := "12345678901"
	fmt.Printf("username: %s, age: %d, address: %s, mobile: %s\n", username, age, address, mobile)
	userLog := fmt.Sprintf("username: %s, age: %d, address: %s, mobile: %s\n", username, age, address, mobile)
	fmt.Println(userLog)

	// 通过string的builder进行字符串拼接 高性能
	fmt.Println("通过string的builder进行字符串拼接 高性能")
	var builder strings.Builder
	builder.WriteString("用户名: ")
	builder.WriteString(username)
	builder.WriteString(", 年龄: ")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(", 地址: ")
	builder.WriteString(address)
	builder.WriteString(", 手机号: ")
	builder.WriteString(mobile)
	re := builder.String()
	fmt.Println(re)

	// 字符串比较
	fmt.Println("字符串比较")
	// 字符串大小比较
	a := "hello"
	b := "bello"
	fmt.Println(a > b)

	// strings用法
	fmt.Println("strings用法")
	// 是否包含 前面的字符串是否包含后面的字符串
	contains := strings.Contains(username, address)
	fmt.Println(contains)
	// 字符串的长度
	usernameBytes := []byte(username)
	length := len(usernameBytes)
	fmt.Println(length)
	// 出现次数 区分大小写
	count := strings.Count(username, "E")
	fmt.Println(count)
	// 分割字符串
	split := strings.Split(username, "_")
	for i, i2 := range split {
		fmt.Println(i, i2)
	}
	// 是否包含前缀 区分大小写
	hasPrefix := strings.HasPrefix(username, "Entic")
	fmt.Println(hasPrefix)
	// 是否包含后缀 区分大小写
	hasSuffix := strings.HasSuffix(username, "Kedaya")
	fmt.Println(hasSuffix)
	// 查找字符串出现的位置
	index1 := strings.Index(username, "浩")
	fmt.Println(index1)
	index2 := strings.LastIndex(username, "浩")
	fmt.Println(index2)
	// 汉字有多个字节这个时候使用rune
	indexRune1 := strings.IndexRune(username, '浩')
	fmt.Println(indexRune1)
	// 查找最后一个重复的中文字符
	lastIndexRune := lastIndexRune(&username, "浩")
	fmt.Println(lastIndexRune)
	// 替换字符串 n=-1的时候替换所有  n=1的时候替换第一个
	replace := strings.Replace(username, "Kedaya", "小火龙", -1)
	fmt.Println(replace)
	// 大小写转换
	toLower := strings.ToLower(username)
	fmt.Println(toLower)
	toUpper := strings.ToUpper(username)
	fmt.Println(toUpper)
	// 去掉特殊字符 Trim只能去掉首尾的字符
	trim := strings.Trim(username, "_")
	fmt.Println(trim)

	fmt.Println("================================")
	sql := make([]string, 0, 5)
	req := "kedaya"
	sql = append(sql, " AND username LIKE %?%", req)
	log.Printf("%v", sql)
	sql = append(sql, " AND username LIKE %?%", req)

}

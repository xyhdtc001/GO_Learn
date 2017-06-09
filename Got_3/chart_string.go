package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// WeekDay 周N
type WeekDay int32

//Sunday iota常量生成器
const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//Char3tString 3章String,rune 字符，byte处理包 strings,bytes...
func Char3tString() {
	s := "abcd"
	r := []rune(s)
	b := []byte(s)
	b[0] = 't'
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
	fmt.Println(string(b))
	fmt.Println(s)
	strings.Contains(s, "a")
	bytes.Contains(b, []byte("a"))
	//bytes.Buffer
}

//StringConverToNumber 字符数字的转换 strconv 包.
func StringConverToNumber() {
	x := 3
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 8))
	//strconv.ParseInt(y, 10, 64)
	var s string
	println("ss", s)
}

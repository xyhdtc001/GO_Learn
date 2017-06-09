package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gotest/Got_1/pkg/tempconv"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

//ChartFlag flag类型
func ChartFlag() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

//NodeMap map
type NodeMap map[int]int

//VarExpressionTest 语法
func VarExpressionTest() {
	//m := make(map[int]int)
	var m NodeMap
	m = make(NodeMap)
	m[1] = 1
	v, ok := m[1]
	fmt.Println(v)
	fmt.Println(ok)
}

//ImportPackage 导入pack测试
func ImportPackage() {
	var va1 tempconv.Interger = 1
	res := va1.Less(tempconv.Interger(2))
	fmt.Println(res)

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf :%v \n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		//c := tempconv.Celsius(t)
		fmt.Printf("%s\n", f)
	}
}

//PrintUse printf 使用技巧
func PrintUse() {
	o := 0666
	fmt.Printf("%d,%[1]o,%#[1]o", o)
}

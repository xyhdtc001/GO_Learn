package main

import (
	"bufio"
	"fmt"
	"os"
)

func _chart_1_3() {
	fmt.Println("Hello Go")
	fmt.Println("argsLen", len(os.Args))
	fmt.Println(os.Args)
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " *** "

	}
	fmt.Println(s)
	fmt.Println("*************st2*************")

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	println(len(counts))
	for line, n := range counts {
		//if n > 1{
		println(line)
		fmt.Printf("%d\t%s\n", n, line)
		//}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "ok" {
			break
		}
		counts[input.Text()]++
	}
}

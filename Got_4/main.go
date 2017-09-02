package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("argsLen", len(os.Args))
	fmt.Println(os.Args)

	_chart4_hashmapUse_()
	//x, _ := strconv.ParseInt(os.Args[1], 10, 8)

	//res := _switch_use(int(x))
	//fmt.Printf("swtich x = %d : %d\n", x, res)

	//array := [2]byte{1, 2}
	//zero(&array)
	//s := []int{0, 1, 2, 3, 4, 5}
	//reverse(s[:2])
	//fmt.Println(s)
	//reverse(s[2:])
	//fmt.Println(s)
	//reverse(s)
	//fmt.Println(s)
}

func zero(ptrArray *[2]byte) {
	*ptrArray = [2]byte{}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

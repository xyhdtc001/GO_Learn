package tempconv

import "fmt"

//Celsius 类型摄氏度
type Celsius float64

//Fahrenheit 类似华氏温度
type Fahrenheit float64

const (
	//AbsoluteZeroC 绝对零度
	AbsoluteZeroC Celsius = -273.15

	//FreezingC 0度
	FreezingC Celsius = 0

	//BoilingC 开水温度
	BoilingC Celsius = 100
)

func (c Celsius) String() string {

	fmt.Println("Celsius string get")
	return fmt.Sprintf("%g ° C", c)
}

func (f Fahrenheit) String() string {
	fmt.Println("Fahrenheit string get")
	return fmt.Sprintf("%g ° F", f)
}

//Interger 整数
type Interger int

//Less 比较大小
func (a Interger) Less(b Interger) bool {
	return a < b
}

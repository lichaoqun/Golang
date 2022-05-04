/*
变量
	变量声明的几种方式
		1. 变量的声明方式 var 变量名 变量类型;
		2. 变量的声明方式并赋值(变量赋值之前一定要先声明) var 变量名 变量类型 = 变量值;
		3. 变量的声明方式并赋值(省去变量类型, 自动推导类型) var 变量名 = 变量值;
		4. 变量的声明方式并赋值(省去变量类型, 自动推导类型) 变量名 := 变量值;
		5. 多重赋值, 只有自动推导类型的声明/初始化的情况可以同时声明/初始化多个类型, 否则只能声明/初始化相同的类型;
	Println 和 Printf
		1. fmt.Println 打印的结果会自动换行, 一段一段的处理; fmt.Printf 不会换行, 自己格式化输出;
		2. fmt.Println 只能使用 fmt.Println("name =", name, ", age = ", age) 方式做格式化, fmt.Printf 可以使用 fmt.Printf("name : %s, age : %d", name, age) 做格式化
	格式化符号
		%c : 以字符方式打印
		%s : 以字符串方式打印
		%t : 以 true / false 方式打印 bool 类型
		%b : 以二进制整数方式打印
		%o : 以八进制整数方式打印
		%d : 以十进制整数方式打印
		%02d : 以十进制整数方式打印, 小于10的数字以 01的方式打印, 如01, 02, 03...
		%x : 以十六进制整数方式打印, 使用 a - f
		%02x : 以十六进制整数方式打印, 使用 a - f, 并且小于10的数字以 01的方式打印, 如01, 02, 03...
		%X : 以十六进制整数方式打印, 使用 A - F
		%02X : 以十六进制整数方式打印, 使用 A - F, 并且小于10的数字以 01的方式打印, 如01, 02, 03...
		%f : 以小数方式打印
		%.2f : 以小数方式打印, 保留两位有效数字打印, 如10.00, 10.01, 10.02...
		%T : 以变量类型方式打印
		%p : 以十六进制整数方式打印地址, 以0x开头
		%v : 使用默认格式输出内置/自定义的类型的值, 如果是自定义的对象并且自定义对象存在String方法, 那么就是使用该类型的String()方式输出值.

	转义字符
		%% : 一个 %
常量
	常量的初始化
	常量的赋值必须在初始化时候赋值, 声明之后就不可以赋值了.
		1. const 常量名 常量类型 = 常量值;
		2. const 常量名(省去变量类型, 自动推导类型) = 常量值;
	变量和常量的区别
		1. 变量是指在程序运行期间可以修改的量, 常量是指在程序运行期间不可以修改的量
		2. 变量声明使用 var 关键字就, 常量声明使用 const 关键字.
多变量/常量的声明
	声明多个变量
		1. 声明多个变量.
		2. 声明多个变量, 然后赋值.
		3. 声明多个变量同时赋值.
		4. 声明多个变量同时赋值(自动推导类型).

	声明多个常量
		1. 声明多个常量同时赋值.
		2. 声明多个常量同时赋值(自动推导类型).

	import 多个包
iota关键字的使用
	1. iota表示常量自动生成器,每一行自动 +1;
	2. iota可以给常量赋值使用
	3. iota 遇到const 自动重置为 0;
	4. 声明时可以只写一个iota;
	5. 如果同iota写在同一行, 那么这一行的值都是一样的;
基本数据类型
	布尔类型
		bool: 8-bit, true ~ false
	数字类型
		整型 (自动类型推导时候默认是 int 类型, 字符型自动类型推导时候默认是 int32 类型)
			无符号型
				uint8: 8-bit, 0 ~ 255
				uint16:16-bit, 0 ~ 65535
				uint32: 32-bit, 0 ~ 4294967295
				uint64: 64-bit, 0 ~ 18446744073709551615
			有符号型
				int8: 8-bit, -128 ~ 127
				int16: 16-bit, -32768 ~ -32767
				int32: 32-bit, -2147483648 ~ 2147483647
				int64: 64-bit, -9223372036854775808 ~ 9223372036854775807
		浮点型(自动类型推导时候默认是 float64 类型)
			float32: 32-bit, 精度 6 位小数
			float64: 64-bit, 精度 15 位小数
	字符串类型
		string
	派生类型
		指针类型
		数组类型
		结构体类型
		Channel类型
		函数类型
		切片类型
		接口类型
		Map类型
	其他类型
		byte: 8-bit, 0 ~ 255 (uint8的别名)
		rune: 32-bit, -2147483648 ~ 2147483647(int32的别名, 表示单个Unicode 字符)
		uintptr: 用于存放指针地址
字符和字符串的区别
	1. 字符使用''(单引号) 而字符串使用 ""(双引号)
	2. 字符由一个字符组成, 字符串由一个或多个字符组成, 并且以结束符 \0 结尾.("abc" 组成方式是 'a''b''c'\0)
类型转换
	1. bool 和 整型 不能相互转换
类型别名
	1. 给某个类型起一个别名
**/
package main

import "fmt"

// - 形数类型的定义, 两个参数, 两个类型
func test1(intputAge int, intName string) (outputName string, outputAge int) {
	outputName = intName
	outputAge = intputAge
	return
}
func test11(intputAge int, intName string) (string, int) {
	return intName, intputAge
}

// - 形数类型的定义, 多个参数, 一个类型
func test2(inputA, inputB int) (outputA, outputB int) {
	outputA = inputA
	outputB = inputB
	return
}
func test22(inputA, inputB int) (int, int) {
	return inputA, inputB
}

// - 形数类型的定义, 一个参数, 多个类型
func test3(inputC string, inputA, inputB int) (outputA, outputB int, outputC string) {
	outputA = inputA
	outputB = inputB
	outputC = inputC
	return
}
func test33(inputC string, inputA, inputB int) (int, int, string) {
	return inputA, inputB, inputC
}

func aboutVar() {
	{ // - 先声明, 但是不赋值, 查看变量的默认值, 数字默认是0, 字符串默认是"";
		var name string
		var age int
		fmt.Printf("name : %s, age : %d \n", name, age)

		// - 多个变量声明, 只能是相同的类型;
		var weight, height float32
		fmt.Printf("name : %f, age : %f \n", weight, height)
	}

	{ // - 先声明后赋值变量的方式
		var name string
		var age int
		name = "李超群1"
		age = 11
		fmt.Printf("name : %s, age : %d \n", name, age)

		// - 多个变量声明,  只能是相同的类型;
		var weight, height float32
		weight = 1.1
		height = 11.1
		fmt.Printf("name : %f, age : %f \n", weight, height)
	}

	{ // - 声明的同时赋值变量
		var name string = "李超群2"
		var age int = 12
		fmt.Printf("name : %s, age : %d \n", name, age)

		// - 多个变量声明, 只能是相同的类型;
		var weight, height float32 = 1.2, 11.2
		fmt.Printf("name : %f, age : %f \n", weight, height)
	}

	{ // - 声明的同时赋值变量(省略变量类型, 自动推导类型)
		var name = "李超群3"
		var age = 13
		fmt.Printf("name : %s, age : %d \n", name, age)

		// - 多个变量声明, 可以是不同的类型
		var weight, height = "1.4", 11.4
		fmt.Printf("name : %s, age : %f \n", weight, height)
	}

	{ // - 声明的同时赋值变量(省略变量类型和var, 自动推导类型)
		name := "李超群4"
		age := 14
		fmt.Printf("name : %s, age : %d \n", name, age)

		// - 多个变量声明
		weight, height := "1.5", 11.5
		fmt.Printf("name : %s, age : %f \n", weight, height)
	}

	{ // - 多重赋值
		name, age := "李超群5", 15
		name1, age1 := name, age

		// - 使用函数的多重赋值
		name2, age2 := test1(20, "李超群")
		fmt.Printf("name : %s, name1 : %s, name2 : %s, age : %d, age1 : %d, age2 : %d \n", name, name1, name2, age, age1, age2)
	}

	{ // - 匿名变量
		name, age := "李超群5", 15

		// - 表示舍弃接收变量age, 只想接收变量name, 这种写法在接收函数多个返回值和遍历字典/数组时候很方便;
		name1, _ := name, age
		_, age2 := test1(20, "李超群")
		fmt.Printf("name : %s, name1 : %s, age : %d, age2 : %d\n", name, name1, age, age2)
	}

	{ // - 交换变量的值
		lcqAge, wxcAge := 20, 10
		fmt.Printf("交换前 : lcqAge : %d  wxcAge: %d\n", lcqAge, wxcAge)

		// - 正常变量交换
		lcqAge, wxcAge = wxcAge, lcqAge

		// - 匿名变量交换
		lcqAge, _ = wxcAge, lcqAge

		fmt.Printf("交换后 : lcqAge : %d  wxcAge: %d\n", lcqAge, wxcAge)
	}

	{ // - 以下代码会有重复定义name变量的错误的错误, 变量赋值之前一定要先声明
		// name = "李超群"
		// var name string
		// name := "李超群"
	}

	{ // - 打印测试测试, fmt.Println 只能将name和age填充在"前边的位置.
		name := "李超群"
		age := 10
		fmt.Println("name = , age = ", name, age)
		fmt.Println("name =", name, ", age = ", age)
	}

	{ // - 格式化符号
		name := "李超群"
		age := 10
		fmt.Printf("name 类型: %T, 值 : %s, age 类型 : %T, 值 : %d", name, name, age, age)
	}
}

func aboutConst() {
	{ // - 定义常量并声明类型
		const age int = 10
		fmt.Println("age : ", age)
	}

	{ // - 定义常量不声明类型
		const age = 10
		fmt.Println("age : ", age)
	}
}

func aboutMutableDefine() {
	{ // - 声明多个变量
		var (
			name   string
			age    int
			height float32
			weight float32
		)
		fmt.Println("name = ", name, ", age = ", age, ", height = ", height, ", weight = ", weight)
	}

	{ // - 声明多个变量, 然后赋值
		var (
			name   string
			age    int
			height float32
			weight float32
		)
		name, age, height, weight = "李超群", 10, 100.0, 101.0
		fmt.Println("name = ", name, ", age = ", age, ", height = ", height, ", weight = ", weight)
	}

	{ // - 声明多个变量同时赋值
		var (
			name   string  = "李超群"
			age    int     = 10
			height float32 = 100.0
			weight float32 = 101.0
		)
		fmt.Println("name = ", name, ", age = ", age, ", height = ", height, ", weight = ", weight)
	}

	{ // - 声明多个变量同时赋值(自动推导类型)
		var (
			name   = "李超群"
			age    = 10
			height = 100.0
			weight = 101.0
		)
		fmt.Println("name = ", name, ", age = ", age, ", height = ", height, ", weight = ", weight)
	}

	{ // - 声明多个常量同时赋值
		const (
			name   string  = "李超群"
			age    int     = 10
			height float32 = 100.0
			weight float32 = 101.0
		)
		fmt.Println("name = ", name, ", age = ", age, ", height = ", height, ", weight = ", weight)
	}

	{ // - 声明多个常量同时赋值(自动推导类型)
		const (
			name   = "李超群"
			age    = 10
			height = 100.0
			weight = 101.0
		)
		fmt.Println("name = ", name, ", age = ", age, ", height = ", height, ", weight = ", weight)
	}

	{ //  -import 多个包
		// import (
		// 	"fmt"
		// 	"go/constant"
		// )
	}
}

func aboutItoa() {
	{ // - iota 每一行自动 +1
		const (
			a = iota
			b = iota
			c = iota
		)
		fmt.Println("a =", a, ", b =", b, ", c =", c)
	}

	{ // - 遇到const 自动置为 0;
		const (
			a = iota
			b = iota
			c = iota
		)
		const d = iota
		fmt.Println("a =", a, ", b =", b, ", c =", c, ", d=", d)
	}

	{ // - 可以声明时候只写一个iota
		const (
			a = iota
			b
			c
		)
		fmt.Println("a =", a, ", b =", b, ", c =", c)
	}

	{ // - 如果多个iota写在同一行, 那么这些常量的值都是相同的;
		const (
			a       = iota
			b, c, d = iota, iota, iota
			e       = iota
		)
		fmt.Println("a =", a, ", b =", b, ", c =", c, ", d =", d, ", e =", e)
	}
}

func aboutType() {
	{ // - int的自动类型推导
		age := 10
		fmt.Printf("age type is %T\n", age)
	}

	{ // - float的自动类型推导
		weight := 10.0
		fmt.Printf("weight type is %T\n", weight)
	}

	{ // - 字符和字符串的区别
		age := 100.0
		fmt.Printf("age type is %.2f %p\n", age, &age)
	}

	{ // - 强制类型转换写法
		ch := 'a'
		var chAsciiVal int = int(ch)
		fmt.Printf("value is : %d\n", chAsciiVal)
	}

	{ // - 类型别名, 定义一个char类型
		type char byte
		var ch char = 'a'
		fmt.Printf("type is %T\n", ch)
	}
}
func main() {
	// - 变量
	aboutVar()

	// - 常量
	aboutConst()

	// - 多个变量/常量的声明
	aboutMutableDefine()

	// - itoa
	aboutItoa()

	// - 基本数据类型
	aboutType()

}

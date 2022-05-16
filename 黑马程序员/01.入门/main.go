/*
	1. go 有且只有一个入口函数 main 函数
	2. "{"左括号必须和函数名同行
	3. go语言不用分号 ";" 结果
	4. 要使用一个包内的功能或者函数, 必须先导入整个包
	5. go语言以包作为管理单位, 每个go语言程序必须先声明包
	6. 每个程序必须有一个main包, 通常程序的入口是 main包中的main函数.
	7. 导入包必须使用, 否则会报错
	8. 变量声明必须使用, 否则会报错
**/
// - 声明包, 每个程序必须有一个main包, 如果没有整个main包, 这时候程序无法运行.
package main

// - 导入包
import "fmt"

// - go 有且只有一个入口函数 main 函数
func main() { //左括号必须和函数名同行
	fmt.Println("Hello World") //Println会自动换行
	fmt.Printf("Hello World")  //go语言结尾可以没有分号
}

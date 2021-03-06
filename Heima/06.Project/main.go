package main

import "fmt"

/* 包管理
import (
	// - 给包起别名, 打印语句变为 myFmt.Println()
	myFmt "fmt"

	// - 忽略包, 这时候即使你导入了这个包没有使用, 也不会报错.
	_ "os"

	// - 调用这个包的函数时候, 不需要 time.Sleep(10), 直接 Sleep(10) 即可
	. "time"
)

**/

/* 工程管理
	同一目录
		同一个目录下go文件的包名必须一样.
		同一目录下调用别的文件的函数, 直接调用即可, 不用引入包名
		同一个包被不允许有同名函数
	不同目录
		不同目录包名可以不一样.
		不同目录下调用别的包的, 必须先导入包, 调用格式 包名.函数名()
		不同包内可以使用同名的函数
		如果包里边的函数、结构体、结构体变量名、的首字母是小写的, 那么别的包是无法调用的, 如果希望这个函数可以被别的包调用, 那么这个函数名的首字母必须大写.
	init函数和main函数
		当 import 一个包的时候, 会执行这个包的 init 函数
		每个包可以有多个init函数
			执行顺序不确定
		每个源文件也可以有多个init函数.
			执行顺序取决函数编写的位置, 自上向下执行.
		import _ "fmt", 可以让这个包执行init函数, 又不会报错(导入不使用会报错)
		init函数和main函数的执行顺序
			导入包的init函数 -> 当前包的init函数 -> 当前包的main函数;
**/
func init() {
	fmt.Println("1")
}
func main() {
}

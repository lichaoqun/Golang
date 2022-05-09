/*
普通函数格式
	func 函数名称(参数列表)(返回值列表){
		函数体
		return 返回值
	}
不定参数的函数
	func 函数名称(形参 ...形参类型)(返回值列表){
		len(形参) : 表示参数的个数;
		形参[i] : 表示访问形参数组中的第i个形参
		函数体
		return 返回值
	}
	不定参数的位置可以不传递参数
	不定参数只能是函数参数列表中的最以后一个参数
	不定参数必须是同一个类型
函数类型(函数指针)
	使用type关键字给一个函数起名,

**/

package main

import "fmt"

func main() {
	func1()
	func2("李超群", "QG", 20, 100.1111, 111.111)
	func3("李超群", "QG", 20, "LCQ", "WXC", "LWT")
	name, _, age := func4("李超群", "QG", 20, "LCQ", "WXC", "LWT")

	inputFuncHandle := func4
	outputFuncHandle, outputName, outputAge := func5(inputFuncHandle, name, age)
	outputName, _, outputAge = outputFuncHandle(outputName, "QG", outputAge, "LCQ", "WXC", "LWT")
	func3(outputName, "QG", outputAge, "LCQ", "WXC", "LWT")

}

// - 无参无返回值0000
func func1() {
	fmt.Println("func1")
}

// - 有参无返回值, 这里的 name, age, weight, height 叫做 形参; 传递进来的参数叫做实参.
func func2(name, nickName string, age int, weight, height float32) {
	fmt.Printf("name: %s, nickName: %s, age: %d, weight: %.2f, height: %.2f;\n", name, nickName, age, weight, height)
}

// - 不定参数
func func3(name, nickName string, age int, friends ...string) {

	for idx, value := range friends { // - 迭代器访问每个形参
		fmt.Println("下标是:", idx, ", 值是:", value)
	}

	fmt.Printf("name : %s, nickName: %s, age: %d, friens: ", name, nickName, age)

	// - 参数全部传递
	func33(friends...)

	// - 传递部分参数
	fmt.Printf("except self is: ")
	func33(friends[1:]...)
}

// - 不定参数的传递
func func33(friends ...string) {

	for i := 0; i < len(friends); i++ { // - for访问每个形参
		fmt.Printf("%s ", friends[i])
	}
	fmt.Printf("\n")
}

// - 有参有返回值
func func4(inputName, inputNickName string, inputAge int, inputFriends ...string) (outputName, outputNickName string, outputAge int) {
	func3(inputName, inputNickName, inputAge, inputFriends[1:]...)
	outputName = "QG"
	outputNickName = "李超群"
	outputAge = 30
	return
}
func func44(inputName, inputNickName string, inputAge int, inputFriends ...string) (outputName, outputNickName string, outputAge int) {
	func3(inputName, inputNickName, inputAge, inputFriends[1:]...)
	outputName = "QG_1"
	outputNickName = "李超群_1"
	outputAge = 31
	return
}

// - 使用函数类型
type funcHandle func(string, string, int, ...string) (string, string, int)

func func5(inputHandle funcHandle, inputName string, inputAge int) (outputHandle funcHandle, outputName string, oututAge int) {
	outputName, _, oututAge = func4("李超群", "QG", 20, "LCQ", "WXC", "LWT")
	outputHandle = func44
	return
}

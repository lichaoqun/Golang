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
匿名函数与闭包
	匿名函数内部可以访问并修改外部的变量. 当匿名函数内部的变量名和外部变量名一致时候, 使用就近原则.
	就近原则
		如果能在本作用域中找到此变量, 就操作此变量
		如果在本作用域中找不到此变量, 外部的变量
	闭包: 就是一个函数捕获了和他在同一个作用域的其他常量和变量, 这表示当闭包被调用的时候, 不管程序在什么地方调用,闭包都能够使用这些常量/变量.(无论捕获的变量是否超出作用域, 所以只要闭包还在使用这些变量/常量, 这些变量/常量就还会存在)
defer 关键字
	函数调用完成后, 执行defer关键字修饰的语句, 程序执行过程中, 执行到defer关键字修饰的语句 就会将defer后边的语句加入到本函数的defer栈中, 先写的的压入栈底,后执行 后写的压入栈顶 先执行
	被defer关键字修饰的语句,只要在函数中被压入了栈中, 那么即使后边的语句出现了crash, 被压入栈中的defer语句也会执行.
defer 和 匿名函数的使用
	如果在一个函数 func1中 使用defer修饰了一个匿名函数, 并且在这个匿名函数有参数, 那么在func1函数执行到defer语句时候, 就已经将参数压入匿名函数栈了, 当 func1 执行完成, 执行defer修饰的语句的时候, 匿名函数栈中的参数是之前压入栈的参数的值, 而不是当前的参数的值
就近原则(局部变量和全局变量同名)
	如果能在本作用域中找到此变量, 就操作此变量
	如果在本作用域中找不到此变量, 就操作全局变量
**/

package main

import "fmt"

func main() {
	{ // - defer 使用
		defer fmt.Println("defer1")
		defer fmt.Println("defer2")
	}

	{ // - 简单的函数调用
		func1()
		func2("李超群", "QG", 20, 100.1111, 111.111)
		func3("李超群", "QG", 20, "LCQ", "WXC", "LWT")
		name, _, age := func4("李超群", "QG", 20, "LCQ", "WXC", "LWT")
		f4 := func4
		f5, outputName, outputAge := func5(f4, name, age)
		outputName, _, outputAge = f5(outputName, "QG", outputAge, "LCQ", "WXC", "LWT")
		fmt.Println(outputName, outputAge)
	}
	{
		// - 匿名函数
		name := "李超群"
		fmt.Print(name)
		funcName := func(inputName, inputNickName string, inputAge int, inputFriends ...string) (outputName, outputNickName string, outputAge int) {
			func3(inputName, inputNickName, inputAge, inputFriends[1:]...)
			// - 当匿名函数内部的变量名和外部变量名一致时候, 在匿名函数内部访问的是匿名函数内部的变量
			outputName = "QG"
			outputNickName = "李超群"
			outputAge = 30
			name = "内部修改外部变量"
			return
		}
		outputName, _, outputAge := funcName("李超群", "QG", 30, "LCQ", "WXC", "LWT")

		// - 定义并调用一个匿名函数
		outputName, _, outputAge = func(inputName, inputNickName string, inputAge int, inputFriends ...string) (outputName, outputNickName string, outputAge int) {
			func3(inputName, inputNickName, inputAge, inputFriends[1:]...)
			// - 当匿名函数内部的变量名和外部变量名一致时候, 在匿名函数内部访问的是匿名函数内部的变量
			outputName = "QG"
			outputNickName = "李超群"
			outputAge = 30
			return
		}(outputName, "QG", outputAge, "LCQ", "WXC", "LWT")
		fmt.Println(outputName, outputAge)
	}

	{ // - 只要闭包还在使用这些变量/常量, 这些变量/常量就还会存在
		f6 := func6()
		fmt.Println(f6())
		fmt.Println(f6())
		fmt.Println(f6())
		fmt.Println(f6())
	}

	{ // - 被defer关键字修饰的语句, crash执行情况

		defer fmt.Println("defer_1")
		defer fmt.Println("defer_2")

		// - 不添加 defer log: defer_2 defer_1
		// - 添加 defer log: defer_4 defer_2 defer_1
		// defer func7(0)
		defer fmt.Println("defer_4")
	}

	{ // - defer配合匿名函数的使用
		a := 10
		b := 20
		// - 在main函数执行完成, 执行defer后边的语句的时候, 匿名函数访问的是外部的变量.
		defer func() {
			fmt.Println("匿名函数没有参数 a :", a, ", b :", b)
		}()
		// - main函数顺序执行, 执行到defer语句的时候, 会将a,b的值压入匿名函数的函数栈中,在main函数执行完成, 执行defer后边的语句的时候, 匿名函数访问的是压入栈的参数.
		defer func(a, b int) {
			fmt.Println("匿名函数有参数 a :", a, ", b :", b)
		}(a, b)
		a = 11
		b = 22
	}

}

// - 无参无返回值
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

// - 匿名函数作为返回值
func func6() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func func7(input int) (output int) {
	output = 100 / input
	return
}

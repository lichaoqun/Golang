/*
指针
	指针的默认值是 nil, 没有NULL一说
	使用操作符 & 取地址.
	在结构体中访问指针变量, 直接使用 . 访问, 不使用 -> 访问.
	指针如果没有具体的指向, 给指针指向的地址赋值是错误的.
	使用new函数申请内存空间, 申请的内存空间不用手动管理.
	new 创建变量
		用new创建变量和普通变量声明语句方式的区别
			用new创建变量和普通变量声明语句方式创建变量没有什么区别, 只是不需要声明一个临时变量的名字, 换而言之new函数类似是一种语法糖，而不是一个新的基础概念
		以下两个函数有这相同的行为
			func newInt() *int {return new(int)}
			func newInt() *int {var dummy int return &dummy }
	变量的生存周期
		在 go 语言中,返回函数中局部变量的地址也是安全的, 因为在调用函数时创建局部变量，在局部变量地址被返回之后依然有效，因为指针依然引用这个变量.
		编译器会自动选择在栈上还是在堆上分配局部变量的存储空间,这个选择并不是由用var还是new声明变量的方式决定的. 因此局部变量可能在函数返回之后依然存在
		对于全局变量来说，它们的生命周期和整个程序的运行周期是一致的。而相比之下，局部变量的生命周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收。

**/
package main

import "fmt"

func aboutPtr() {
	{ // - 指针
		// - 一级指针和二级指针的简单使用
		value := 100
		ptr1 := &value
		var ptr11 **int = &ptr1
		fmt.Printf("%p 存储的值是 %p, %p 存储的值是 %d\n", ptr11, *ptr11, *ptr11, **ptr11)

		// - 指针默认值是nil;
		var ptr2 *int
		fmt.Printf("指针默认值是 : %p\n", ptr2)

		// - 指针如果没有具体的指向, 给指针指向的地址赋值是错误的
		// - 以下操作会报错
		// *f2 = 100

		// - 使用new函数申请内存空间
		ptr3 := new(int)
		*ptr3 = 100
		fmt.Printf("%p 存储的值是 %d\n", ptr3, *ptr3)

		// - 值传递和指针传递
		value2 := 20
		swapValue(value, value2)
		fmt.Println("value is :", value, ", value2 is :", value2)

		// - 指针传递
		swapPtr(&value, &value2)
		fmt.Println("value is :", value, ", value2 is :", value2)
	}

	{ // - go 语言中,返回函数中局部变量的地址也是安全的, 以下代码在调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。所以以下代码可以正常执行
		var p = f()
		fmt.Println(p)
	}
}

// - 值传递
func swapValue(value1, value2 int) {
	value1, value2 = value2, value1
}

// - 指针传递
func swapPtr(value1, value2 *int) {
	*value1, *value2 = *value2, *value1
}


func () *int {
	v := 1
	return &v
}

// - fTest函数里的x变量必须在堆上分配，因为它在函数退出后依然可以通过包一级的global变量找到，虽然它是在函数内部定义的；用Go语言的术语说，这个x局部变量从函数f中逃逸了。
// - gTest函数返回时，变量*y将是不可达的，也就是说可以马上被回收的。因此，*y并没有从函数g中逃逸，编译器可以选择在栈上分配*y的存储空间
var global *int
func fTest() {
	var x int
	x = 1
	global = &x
}

func gTest() {
	y := new(int)
	*y = 1
}

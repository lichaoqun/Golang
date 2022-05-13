/*
指针
	指针的默认值是 nil, 没有NULL一说
	使用操作符 & 取地址.
	在结构体中访问指针变量, 直接使用 . 访问, 不使用 -> 访问.
	指针如果没有具体的指向, 给指针指向的地址赋值是错误的.
	使用new函数申请内存空间, 申请的内存空间不用手动管理.
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
}

// - 值传递
func swapValue(value1, value2 int) {
	value1, value2 = value2, value1
}

// - 指针传递
func swapPtr(value1, value2 *int) {
	*value1, *value2 = *value2, *value1
}

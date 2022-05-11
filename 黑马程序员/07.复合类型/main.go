/*
指针
	指针的默认值是 nil, 没有NULL一说
	使用操作符 & 取地址.
	在结构体中访问指针变量, 直接使用 . 访问, 不使用 -> 访问.
	指针如果没有具体的指向, 给指针指向的地址赋值是错误的.
	使用new函数申请内存空间, 申请的内存空间不用手动管理.
数组
	同一类型的集合
	声明数组时候, 数组长度不能用变量, 只能用常量.
	访问数组元素时候, 数组的下标可以用变量, 也可以用常量.
	数组作为函数参数传递
	数组作为函数参数, 是值传递

随机数函数
	函数中的种子只需要设置一次
	如果种子参数一样, 每次程序运行产生随机数都是相同的.

slice

**/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
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

	{ // - 数组的声明和使用
		// - 声明数组, 注意数组的长度必须是一个常量, 不能是变量
		// - 下边写法会报错
		// var varLen = 3
		// var array2 [varLen]int

		const count = 3
		var ages [count]int

		// - for 遍历数组
		for idx := 0; idx < len(ages); idx++ {
			ages[idx] = idx + 1
		}

		// - 迭代器遍历数组
		for idx, age := range ages {
			fmt.Println("index is :", idx, ", value is :", age)
		}

		// - 数组的初始化
		// - 普通初始化方案
		var nums1 [count]int = [count]int{1, 2, 3}
		fmt.Println(nums1)

		// - 自动类型推导
		var nums2 = [count]int{1, 2, 3} // - 自动类型推导
		fmt.Println(nums2)

		// - 自动类型推导
		nums3 := [count]int{1, 2, 3}
		fmt.Println(nums3)

		// - 默认填充元素
		nums4 := [count]int{1}
		fmt.Println(nums4)

		// - 指定某个元素初始化
		nums5 := [count]int{0: 3, 2: 1}
		fmt.Println(nums5)

		// - 二维数组
		// - 三行四列的二位素组初始化
		nums6 := [3][4]int{
			{11, 21, 31, 41},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
		}
		fmt.Println(nums6)

		// - 默认填充元素和指定某个元素初始化
		var nums7 [3][4]int = [3][4]int{0: {11, 21}, 2: {0: 31, 2: 33}}
		fmt.Println(nums7)
	}

	{ // - 随机数 这里因为种子是确定的, 所以程序每次启动得到的随机数都是相同的, 如果不希望每次的随机数都相同, 可以使用时间作为随机数种子.
		rand.Seed(99)
		for i := 0; i < 4; i++ {
			// - 随便产生一个随机数
			fmt.Println("产生的随机数为:", rand.Int())

			// - 产生的随机数限制在某个区间
			fmt.Println("产生的随机数为:", rand.Intn(100))
		}
	}

	{ // - 数组和数组指针作为参数
		array := [3]int{0, 1, 2}
		fmt.Println("初始状态: ", array)

		// - 数组作为参数是值传递, 只是将数组元素复制后作为参数参数.
		modify1(array)
		fmt.Println("调用 数组作为参数 后: ", array)
		// - 数组指针作为参数指针传递
		modify2(&array)
		fmt.Println("调用 数组指针作为参数 后: ", array)
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

// - 数组作为函数参数传递
func modify1(array [3]int) {
	array[0] = 2
	array[1] = 1
	array[2] = 0
}

func modify2(arrayPtr *[3]int) {
	(*arrayPtr)[0] = 2
	(*arrayPtr)[1] = 1
	(*arrayPtr)[2] = 0
}

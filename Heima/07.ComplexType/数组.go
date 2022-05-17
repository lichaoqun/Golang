/*
数组
	同一类型的集合
	声明数组时候, 数组长度不能用变量, 只能用常量.
	访问数组元素时候, 数组的下标可以用变量, 也可以用常量.
	数组作为函数参数传递
	数组作为函数参数, 是值传递
**/
package main

import "fmt"

func aboutArray() {
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

		{ // - 数组和数组指针作为参数
			array := [3]int{0, 1, 2}
			fmt.Println("初始状态: ", array)

			// - 数组作为参数是值传递, 只是将数组元素复制后作为参数参数.
			modifyArray1(array)
			fmt.Println("调用 数组作为参数 后: ", array)

			// - 数组指针作为参数指针传递
			modifyArray2(&array)
			fmt.Println("调用 数组指针作为参数 后: ", array)
		}
	}

	{ // - 二维数组
		// - 三行四列的二位素组初始化
		nums1 := [3][4]int{
			{11, 21, 31, 41},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
		}
		fmt.Println(nums1)

		// - 默认填充元素和指定某个元素初始化
		var nums2 [3][4]int = [3][4]int{0: {11, 21}, 2: {0: 31, 2: 33}}
		fmt.Println(nums2)
	}
}

// - 数组作为函数参数传递
func modifyArray1(array [3]int) {
	array[0] = 2
	array[1] = 1
	array[2] = 0
}

func modifyArray2(arrayPtr *[3]int) {
	(*arrayPtr)[0] = 2
	(*arrayPtr)[1] = 1
	(*arrayPtr)[2] = 0
}

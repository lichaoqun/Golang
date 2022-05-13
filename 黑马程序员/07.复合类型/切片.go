/*
slice
	slice 和 底层的数组使用相同的一块内存.
	数组初始化和声明时候需要指定长度, 而slice 不用指定长度
	写法
		s1 := [...]int{1, 2, 3, 0, 0}
		s1 := []int{1, 2, 3, 0, 0}
		s1 := make([]int, 3 ,5) : make(切片类型, 切片长度, 切片容量)
	元素的访问
		元素的访问方式和数组相同
	使用 slice 截取数组的指定元素
		slice := array[low:high:max]
		low: low 索引的起点
		high: 长度截止的索引(不包括此索引); 即长度范围是 [low, high),  此时 slice 的长度 = high - low
		max: 容量截止的索引(不包括此索引); 即容量范围是 [low, max),  此时 slice 的容量 = max - low
		长度 : 表示已填充内容的长度  len(slice) = high - low
		容量 : 表示已填充内容的长度 + 未填充内容的长度; cap(slice) = max - low
		使用append函数在slice尾部追加元素, 当需要扩容的时候, 就会将容量扩大位原来的两倍;
	区别
		长度和容量
			数组长度和容量不可变.
			切片长度和容量可以不固定.
		作为函数参数
			数组作为函数参数是值拷贝.
			切片作为函数参数是引用.
**/
package main

import (
	"fmt"
	"math/rand"
)

func aboutSlice() {
	{ // - slice
		// - 切片初始化方式并类型推导
		s1 := [...]int{1, 2, 3, 0, 0}
		fmt.Println("s1 是:", s1, "长度是:", len(s1), "容量是:", cap(s1))

		// - 切片初始化方式并类型推导
		var s2 = [...]int{1, 2, 3, 0, 0}
		fmt.Println("s2 是:", s2, "长度是:", len(s2), "容量是:", cap(s2))

		// - 切片初始化方式并类型推导
		var s3 = []int{1, 2, 3, 0, 0}
		fmt.Println("s3 是:", s3, "长度是:", len(s3), "容量是:", cap(s3))

		// - 切片的访问方式和数组相同
		fmt.Println("s3 是:", s3, "第2个元素是:", s3[2], "长度是:", len(s3), "容量是:", cap(s3))

		// - 使用make函数
		var s4 = make([]int, 3, 5)
		fmt.Println("s4 是:", s4, "长度是:", len(s4), "容量是:", cap(s4))

		// - 长度和容量都是5
		s5 := make([]int, 5)
		fmt.Println("s5 是:", s5, "长度是:", len(s5), "容量是:", cap(s5))
	}

	{ // - 切片的截取操作

		array := [5]int{1, 2, 3, 4, 5}

		// - 指定起始位置和长度截止位置和容量截止位置.
		nums1 := array[2:3:4]
		fmt.Println("nums1 是:", nums1, "长度是:", len(nums1), "容量是:", cap(nums1))

		// - 指定起始位置 长度和容量 = len(array) - low.
		nums2 := array[2:]
		fmt.Println("nums2 是:", nums2, "长度是:", len(nums2), "容量是:", cap(nums2))

		// - 指定长度, 起始位置从0开始, 容量是数组长度.
		nums3 := array[:3]
		fmt.Println("nums3 是:", nums3, "长度是:", len(nums3), "容量是:", cap(nums3))

		// - 长度和容量都使用数组长度, 起始位置从0开始
		nums4 := array[:]
		fmt.Println("nums4 是:", nums4, "长度是:", len(nums4), "容量是:", cap(nums4))
	}

	{ // - 切片对底层数组的影响
		array := [5]int{1, 2, 3, 4, 5}

		// - 数组的操作影响了slice, slice的操作影响了数组, 说明两个对象共用一块内存.
		nums1 := array[1:4]
		nums1[0] = 100
		nums1[1] = 200
		array[3] = 300
		fmt.Println("nums1 是:", nums1, "array 是:", array)
	}

	{ // - append 函数
		s1 := []int{1, 2, 3, 4, 5}
		s2 := append(s1, 6)
		fmt.Println("s1 是:", s1, ",  s2 是:", s2)

		// - 扩容验证
		s3 := make([]int, 0, 1)
		oldCap := cap(s3)
		for i := 0; i < 20; i++ {
			s3 = append(s3, 1)
			if newCap := cap(s3); newCap > oldCap {
				fmt.Println("old cap is :", oldCap, ", new cap is :", newCap)
				oldCap = newCap
			}
		}
	}

	{ // - copy 函数, 将src中的元素拷贝到dst中, 最多拷贝src的前len(dst)长度到dst中.
		srcSlice1 := []int{1, 2}
		desSlice := []int{6, 7, 8, 9}

		// - len(src) < len(dst)
		copy(desSlice, srcSlice1)
		fmt.Println("src :", srcSlice1, ", dst :", desSlice)

		// - len(src) > len(dst)
		copy(desSlice, srcSlice1)
		fmt.Println("src :", srcSlice1, ", dst :", desSlice)
	}

	{ // - slice 作为函数参数
		slice := make([]int, 5)
		modifySlice(slice)
		fmt.Println("src :", slice)
	}
}

// - slice 作为函数参数
func modifySlice(slice []int) {
	rand.Seed(9)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100)
	}
}

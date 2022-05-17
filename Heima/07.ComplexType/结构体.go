/*
	对于结构体指针和普通变量指针, 访问方式都是一致的, 都是直接通过 . 访问成员;
	结构体比较, 当连个结构体的成员完全相同, 可以相互赋值.
	结构体作为参数是值传递,可以使用指针传递
**/
package main

import "fmt"

// - 使用枚举
type sexType int32

const (
	sexTypeMr   sexType = 1
	sexTypeMiss sexType = 2
)

type Student struct {
	id      int
	name    string
	sex     sexType
	age     int
	address string
	height  *int
}

func aboutStruct() {
	{ // - 初始化
		heiht := 130

		// - 结构体初始化
		var s1 Student = Student{10001, "LCQ1", sexTypeMr, 10, "内蒙古1", new(int)}
		fmt.Println("s1 is", s1)

		// - 自动推导类型
		var s2 = Student{20002, "LCQ2", sexTypeMiss, 20, "内蒙古2", &heiht}
		fmt.Println("s2 is", s2)

		// - 自动推导类型
		s3 := Student{30003, "LCQ3", sexTypeMr, 30, "内蒙古3", &heiht}
		fmt.Println("s3 is", s3)

		// - 指定成员初始化, 未指定的使用默认值
		s4 := Student{name: "LCQ4", age: 40}
		fmt.Println("s4 is", s4)

		// - 使用new函数初始化
		s5 := new(Student)
		s5.name = "LCQ5"
		fmt.Println("s5 is", s5)

		// - 先定义在赋值
		var s6 *Student
		s6 = &s4
		fmt.Println("s6 is", s6)

		// - 先定义在赋值
		var s7 Student
		s7 = s4
		fmt.Println("s7 is", s7)
	}

	{ // - 结构体指针
		// - 结构体指针
		s1 := &Student{name: "LCQ1", age: 10}
		fmt.Println("s1 is", *s1)

		// - 结构体指针
		var s2 = &Student{name: "LCQ2", age: 20}
		fmt.Println("s2 is", *s2)

		// - 结构体指针
		var s3 *Student = &Student{name: "LCQ3", age: 30}
		fmt.Println("s3 is", *s3)
	}

	{ // - 结构体成员的访问

		heiht := 130

		// - 普通结构体类型的成员的访问
		s1 := Student{10001, "LCQ1", sexTypeMiss, 10, "内蒙古1", &heiht}
		fmt.Println("s1. name is", s1.name)

		// - 普通结构体类型的成员的赋值
		s1.name = "L2"
		fmt.Println("s1. name is", s1.name)

		// - 结构体指针类型的成员的访问
		s2 := &Student{20002, "LCQ2", sexTypeMiss, 20, "内蒙古2", &heiht}
		fmt.Println("s2.age is", s2.age)

		// - 结构体指针类型的成员的赋值
		s2.age = 22
		fmt.Println("s2.age is", s2.age)
	}

	{ // - 结构体作为参数
		// - 结构体作为参数
		heiht1 := 130
		s1 := Student{10001, "LCQ1", sexTypeMr, 10, "内蒙古1", &heiht1}
		fmt.Println("s1 is", s1, "s1.height is", *(s1.height))
		modifyStruct(s1)
		fmt.Println("s1 is", s1, "s1.height is", *(s1.height), "var height is", heiht1)

		// - 结构体指针作为参数
		heiht2 := 130
		s2 := Student{10001, "LCQ1", sexTypeMr, 10, "内蒙古1", &heiht2}
		fmt.Println("s2 is", s2, "s2.height is", *(s2.height))

		// - 结构体指针作为参数
		modifyStructPtr(&s2)
		fmt.Println("s2 is", s2, "s2.height is", *(s2.height), "var height is", heiht2)
	}
}

func modifyStruct(s Student) {
	s.name = "LCQ"
	*(s.height) = 1
}

func modifyStructPtr(s *Student) {
	s.name = "LCQ"
	*(s.height) = 1
}

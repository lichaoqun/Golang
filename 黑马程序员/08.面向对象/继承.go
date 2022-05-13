/*
继承
	go 的面向对象 还是用结构体
	继承只是在当前结构体内引入其他结构体
	子类和父类拥有同名变量
		当子类和父类拥有同名变量时候,各自使用各自的变量, 不会重写
		修改同名变量时候, 会使用就近原则
			如果能在本作用域中找到此成员, 就操作此成员
			如果在本作用域中找不到此成员, 就操作继承的字段
**/
package main

import "fmt"

type Person struct {
	name   string
	age    int
	height float32
}

type Student struct {

	// - 使用匿名结构体做继承
	Person

	id int

	nickName string
}

type Coder struct {

	// - 使用匿名结构体做继承
	Person

	// - 工龄
	age int
}

type Child struct {

	// - 使用匿名结构体做继承
	*Person

	// - 压岁钱
	int
}

func abuoutJC() {
	{ // - 初始化
		// - 直接初始化
		s1 := Student{Person{"LCQ1", 1, 1.1}, 111, "QG1"}
		fmt.Println("s1 is", s1)

		//  - 使用变量初始化
		p1 := Person{"LCQ2", 2, 2.2}
		id := 222
		s2 := Student{p1, id, "QG2"}
		fmt.Println("s2 is", s2)

		// - 指定成员初始化
		s3 := Student{Person: Person{name: "LCQ3", height: 3.3}, nickName: "QG3"}
		fmt.Println("s3 is", s3)
	}

	{
		// - 成员访问
		s1 := Student{Person{"LCQ1", 1, 1.1}, 111, "QG1"}
		fmt.Println("s1 is", s1, ", s1.name is", s1.name, ", s1.id is", s1.id, "s1.person is", s1.Person)

		// - 成员赋值
		s1.Person = Person{name: "LCQ2", height: 2.2}
		s1.Person.age = 2
		s1.id = 222
		fmt.Println("s1 is", s1, ", s1.name is", s1.name, ", s1.id is", s1.id, "s1.person is", s1.Person)
	}

	{
		// - 同名变量
		s1 := Coder{Person{"LCQ1", 1, 1.1}, 11}
		fmt.Println("s1 is", s1)

		// - 操作子类的同名变量的成员
		s1.age = 22
		fmt.Println("s1.age is", s1.age)

		// - 操作父类类的同名变量的成员
		s1.Person.age = 2
		fmt.Println("s1.Person.age is", s1.Person.age)
	}

	{
		// - 指针类型匿名字段
		s1 := Child{&Person{"LCQ1", 1, 1.1}, 11}
		fmt.Println("s1.name is", s1.name, ", s1.age is", s1.Person.age, ", s1.height is", s1.Person.height, "s1.money is", s1.int)

		s1.int = 111111
		s1.Person.name = "LCQ"
		s1.age = 22
		fmt.Println("s1.name is", s1.name, ", s1.age is", s1.Person.age, ", s1.height is", s1.Person.height, "s1.money is", s1.int)
	}
}

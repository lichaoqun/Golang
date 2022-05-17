/*
方法
	方法是给 "自定义类型" 添加的
	不允许给简单类型添加方法, 不能给int添加方法, 所以如果想给int添加方法 就 type Integer int, 这样就有了自定义类型 Integer, 就可以添加方法了.
	方法的调用是值传递, 如果想让方法也跟着改, 需要使用指针 / 返回值
	可以给自定义类型添加方法, 但是不能给自定义的指针类型添加方法

		可以编译通过 ,因为 Cat 是自定义类型
			type Cat struct {}
			func (c *Cat) init3() (result *Cat) {}

		无法编译通过, 因为NSCat 是自定义的指针类型
			type NSCat *Cat
			func (c NSCat) init3() (result NSCat) {}
	方法不支持重载, 只要接受者类型不一样, 即使方法同名,也是不同的方法
	指针变量和普通变量的方法集
		指针变量和普通变量可以互相调用对方的方法, go会将当前的变量转化成对应的变量类型, 然后再调用方法
	方法重写遵守就近原则
		现在本作用域中找方法, 如果找不到就调用继承的方法
	方法表达式中 需要使用明确的类型, go不会自动做指针和对象的转换, 需要自己做类型转换;

**/
package main

import "fmt"

func aboutMXDX() {
	fmt.Println("ab")
	{ // - 给int添加方法
		var a Integer = 1
		var b Integer = 2
		var c Integer = a.add(b)
		fmt.Println("a is", a, ", b is", b, ", c is", c)
	}

	{
		// - 结构体添加方法
		var c1 Cat = Cat{"猫", 1, 1.1}
		c1.init1()
		fmt.Println("cat is", c1)

		// - 结构体添加方法, new
		c21 := new(Cat)
		c21.init2()
		fmt.Println("cat is", c21)

		// - 结构体添加方法, 取地址
		var c Cat
		c22 := &c
		c22.init2()
		fmt.Println("cat is", c22, ", c is", c)

		// - 结构体添加方法并且有返回值, new
		c31 := new(Cat).init3()
		fmt.Println("cat is", c31)

		// - 结构体添加方法并且有返回值, 取地址
		c32 := (&c).init3()
		fmt.Println("cat is", c32, ", c is", c)
	}

	{ // - 指针变量和普通变量的方法集
		var c1 Cat = Cat{"猫", 1, 1.1}

		// - 普通变量调用两个方法
		c1.init1()
		c1.init2()

		// - 指针变量调用两个方法
		c2 := &c1
		c2.init1()
		c2.init2()
	}

	{ // - 方法继承
		var c1 Cat = Cat{"猫", 1, 1.1}
		var c2 Garfield = Garfield{c1}

		// - 父类子类都实现
		c1.eat()
		c2.eat()

		// - 父类子类都实现, 但是使用子类调用父类的实现
		c2.Cat.eat()

		// - 父类子类都实现, 子类中调用父类的方法
		c1.run()
		c2.run()

		// - 父类实现, 子类未实现就会自动调用父类的方法
		c1.play()
		c2.play()
	}

	{ // - 方法作为方法参数
		var c1 Cat = Cat{"猫", 1, 1.1}
		c1.customMethod(c1.customMethodParams)

		// - 方法值
		func1 := c1.customMethodParams
		func1(int(c1.age))

		// - 方法表达式, 需要 *Cat, 使用 Cat类型转换
		func2 := (*Cat).init2
		func2(&c1)
		fmt.Println(c1)

		// - 方法表达式, 需要 Cat, 使用 *Cat类型转换
		c2 := &c1
		func3 := (Cat).init1
		func3(*c2)
		fmt.Println(*c2)

		// - 以下代码会报错, 这时候不会自动做类型转换, 需要自己做类型转换;
		// func2(c1)
		// func3(c2)

	}
}

// - 给int加方法, 如果直接给int添加方法会报错, 不允许给简单类型添加方法, 所以给int重新定义一个type, 然后给这个type添加方法

type Integer int

func (a Integer) add(b Integer) (c Integer) {
	a = a + b
	return a
}

// - 以下会报错重复定义重名函数, 方法没有重载
// func (a Integer) add(b1, b2 Integer) (c Integer)

// - 给结构添加方法
type Cat struct {
	name   string
	age    Integer
	height float32
}

// - 值传递直接修改
func (c Cat) init1() {
	c.name = "小猫"

	var a Integer = 1
	var b Integer = 2
	c.age = a.add(b)

	c.height = 1.7
}

// - 使用指针
func (c *Cat) init2() {
	c.name = "init2"

	var a Integer = 1
	var b Integer = 2
	c.age = a.add(b)

	c.height = 2.2
}

// - 使用指针
func (c *Cat) init3() (result *Cat) {
	c.name = "init3"

	var a Integer = 1
	var b Integer = 2
	c.age = a.add(b)

	c.height = 3.3

	result = c
	return
}

// - 一下无法编译通过, 可以给自定义类型添加方法, 但是自定义类型不能是指针类型
// - 使用指针
// type NSCat *Cat
// func (c NSCat) init3() (result NSCat) {}

func (c Cat) eat() {
	fmt.Println("Cat eat...")
}

func (c Cat) run() {
	fmt.Println("Cat run...")
}

func (c Cat) play() {
	fmt.Println("Cat play...")
}

// - 加菲猫
type Garfield struct {
	Cat
}

func (c Garfield) eat() {
	fmt.Println("Garfield eat...")
}

func (c Garfield) run() {
	c.Cat.run()
	fmt.Println("Garfield run...")
}

// - 方法作为方法函数参数
func (c Cat) customMethodParams(x int) int {
	fmt.Println("customMethodParams")
	return int(c.age) + x
}

func (c Cat) customMethod(customMethodInput func(int) int) {
	fmt.Println("customMethodInput 前")
	result := customMethodInput(int(c.age))
	fmt.Println("customMethodInput 后", ", result is", result)
}

/*
接口
	使用接口作为函数参数
		使用接口作为函数参数, 函数参数必须是接口对象, 不能是接口对象指针 如:func runTest(r Runner) {}.
	接口是可以被继承的, 类似结构体的继承
	接口的多态
		使用接口接收的变量类型必须是实现这个接口时候绑定的变量类型
		已知  type Runner interface { run() }
			如果你的接口中使用 func (tmp Car) run() 那么你需要这么写 var r Runner = Car{}
			如果你的接口中使用 func (tmp *Bird) run(), 那么你需要这么写 var b AnimalRunner = &Bird{}
	空接口
		空接口可以当做泛型使用 类似 void* id object
	获取类型
		使用 data.(type) 获取是对象类型, 使用 data.(int)判断对象是不是 int类型
	结构体中可以使用匿名接口

**/
package main

import "fmt"

type Runner interface {
	run()
}

type AnimalRunner interface {
	Runner
	eat()
}

type Dog struct {
	title    string
	runUse   string
	legCount int
}

func (tmp *Dog) run() {
	fmt.Println(tmp.title, "使用", tmp.legCount, "条", tmp.runUse, "跑")
}

func (tmp *Dog) eat() {
	fmt.Println(tmp.title, "eat")
}

type Car struct {
	title      string
	runUse     string
	wheelCount int
}

func (tmp Car) run() {
	fmt.Println(tmp.title, "使用", tmp.wheelCount, "个", tmp.runUse, "跑")
}

type Bird struct {
	title     string
	runUse    string
	wingCount int
}

// - 结构体中使用匿名接口
type Anchor struct {
	AnimalRunner
	livePlatform string
}

func (tmp *Bird) run() {
	fmt.Println(tmp.title, "使用", tmp.wingCount, "个", tmp.runUse, "飞")
}

func (tmp *Bird) eat() {
	fmt.Println(tmp.title, "eat")
}

// - 以下函数会报错, 因为接口作为函数参数, 必须是对象, 不能是对象指针
// func runTest(r *Runner) {
// 	(*r).run()
// }

func runTest(r Runner) {
	r.run()
}

func personRunnerTest(p AnimalRunner) {
	p.run()
	p.eat()
}

func aboutJK() {
	p := &Dog{"小狗", "腿", 4}
	c := Car{"小汽车", "轱辘", 4}
	b := Bird{"小鸟", "翅膀", 2}

	{ // - 使用结果作为函数参数
		runTest(p)
		runTest(c)
		runTest(&b)
	}

	{ // - 使用接口作为数组元素
		slice := make([]Runner, 3)
		slice[0] = p
		slice[1] = c
		slice[2] = &b

		for _, runner := range slice {
			runTest(runner)
		}
	}

	{ // - 接口继承
		personRunnerTest(p)
		personRunnerTest(&b)
	}

	{ // - 接口多态
		var r Runner = Car{"小汽车", "轱辘", 4}
		var b AnimalRunner = &Bird{"小鸟", "翅膀", 2}

		runTest(r)
		runTest(b)

		personRunnerTest(b)
		//- 下边会报错, 多态必须是父类指向子类
		// personRunnerTest(r)
	}

	{ // - 结构体中使用匿名接口
		a := Anchor{&b, "企鹅体育"}

		// - 调用接口的方法
		a.AnimalRunner.run()

		fmt.Println(a, "直播平台是", a.livePlatform)
	}

	{ // - 空接口
		//-  空接口类型的声明
		var r Runner = Car{"小汽车", "轱辘", 4}
		var i1 interface{} = 1
		var i2 interface{} = 2.2
		var i3 interface{} = '3'
		var i4 interface{} = "4"
		var i5 interface{} = false
		var i6 interface{} = p
		var i7 interface{} = c
		var i8 interface{} = r

		// - 使用数组的空接口类型
		slice := make([]interface{}, 0)
		slice = append(slice, i1, i2, i3, i4, i5, i6, i7, i8)

		// - 打印
		for _, interfaceValue := range slice {
			fmt.Printf("type is %T, value is %v\n", interfaceValue, interfaceValue)
		}

		// - 类型判断
		logType1(i1, i2, i3, i4, i5, i6, i7, i8)
		logType2(i1, i2, i3, i4, i5, i6, i7, i8)
	}
}

// - 任意类型的可变参数作为函数参数
func logType1(arg ...interface{}) {
	// - 类型查询
	for idx, data := range arg {
		if value, isType := data.(int); isType {
			fmt.Println("第", idx, "的类型是int", ", value is", value)
		}
		if value, isType := data.(*Dog); isType {
			fmt.Println("第", idx, "*Dog", ", value is", value)
		}
		if value, isType := data.(Car); isType {
			fmt.Println("第", idx, "Car", ", value is", value)
		}
		if value, isType := data.(AnimalRunner); isType {
			fmt.Println("第", idx, "AnimalRunner", ", value is", value)
		}
	}
}

func logType2(arg ...interface{}) {
	for idx, data := range arg {
		switch data.(type) {
		case int:
			fmt.Println("第", idx, "是基本数据类型 int, value is", data)
		case float64:
			fmt.Println("第", idx, "是基本数据类型 float64, value is", data)
		case int32:
			fmt.Println("第", idx, "是基本数据类型 int32, value is", data)
		case string:
			fmt.Println("第", idx, "是基本数据类型 string, value is", data)
		case bool:
			fmt.Println("第", idx, "是基本数据类型 bool, value is", data)
		default:
			fmt.Println("第", idx, "是自定义数据类型, value is", data)
		}
	}
}

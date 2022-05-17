/*
流程控制
	1. 选择结构
		if else
			支持一个初始化语句, 初始化语句和判断语句以 ; 分隔;
		switch
			支持一个初始化语句, 初始化语句和变量语句以 ; 分隔;
			不用写break;
			fallthrough 关键字表示不跳出switch语句, 如果case通过,下边的都无条件执行;
			case后边接条件(注意此时switch后边是没有变量的);
	2. 循环结构
		for
			第一个语句, 只有在第一次进入for循环时候执行;
				如果直接使用外部变量可以省略;
			第二个语句, 每次执行for循环{}内的语句前都会先判断;
				判断条件可以省略
			第三个语句, 每次正常执行完for循环{}内的语句后都会执行一次;
				continue 跳过本次循环会执行第三个语句;
				break 跳出循环不会执行第三个语句;
				可以内部使用goto跳出循环, 跳出循环不会执行第三个语句;
				可以省略
		goto
**/
package main

import "fmt"

func main() {
	{ // - 选择结构 if使用
		s := true
		if s {
			fmt.Println("s == true")
		}
	}

	{ // - 选择结构 if判断的同时去做初始化的操作
		if s := 4; s >= 3 {
			fmt.Println("s >= 3")
		}
	}

	{ // - 选择结构 if else
		if s := 2; s >= 3 {
			fmt.Println("s >= 3")
		} else {
			fmt.Println("s < 3")
		}
	}

	{ // - 选择结构 if else if else
		if s := 3; s > 3 {
			fmt.Println("s > 3")
		} else if s == 3 {
			fmt.Println("s == 3")
		} else {
			fmt.Println("s < 3")
		}
	}

	{ // - 选择结构 switch, 在 go 语言中, 是不用写break的;
		s := 1
		switch s {
		case 1:
			fmt.Println("s == 1")
		case 2:
			fmt.Println("s == 2")
		case 3:
			fmt.Println("s == 3")
		default:
			fmt.Println("都不是哦")
		}
	}

	{ // - 选择结构 switch, 使用fallthrough, fallthrough 关键字表示不跳出switch语句, 如果case通过, 下边的都无条件执行;
		s := 2
		switch s {
		case 1:
			fmt.Println("s == 1")
			fallthrough
		case 2:
			fmt.Println("s == 2")
			fallthrough
		case 3:
			fmt.Println("s == 3")
			fallthrough
		default:
			fmt.Println("都不是哦")
		}
	}

	{ // - 选择结构 switch, 同时初始化
		switch s := 2; s {
		case 1:
			fmt.Println("s == 1")
			fallthrough
		case 2:
			fmt.Println("s == 2")
			fallthrough
		case 3:
			fmt.Println("s == 3")
			fallthrough
		default:
			fmt.Println("都不是哦")
		}
	}

	{ // - 选择结构 switch, case后边方条件
		switch s := 2; {
		case s < 1:
			fmt.Println("s < 1")
		case s < 3:
			fmt.Println("s < 3")
		case s < 5:
			fmt.Println("s < 5")
		default:
			fmt.Println("都不是哦")
		}
	}

	{ // - 选择结构 switch, case后边方条件, 同时使用 fallthrough 关键字;
		switch s := 2; {
		case s < 1:
			fmt.Println("s < 1")
			fallthrough
		case s < 3:
			fmt.Println("s < 3")
			fallthrough
		case s < 5:
			fmt.Println("s < 5")
			fallthrough
		default:
			fmt.Println("都不是哦")
		}
	}

	{ // - for 循环
		for idx := 0; idx < 3; idx++ {
			fmt.Println("idx = ", idx)
		}
	}

	{ // - range 迭代
		name := "Li.CQ"
		for idx, value := range name {
			fmt.Printf("idx = %d, value = %c\n", idx, value)
		}

		// - 使用匿名变量
		for _, value := range name {
			fmt.Printf("value = %c\n", value)
		}
	}

	{ // - goto 语句的使用

		if value := 1; value < 2 {
			goto end
		}
		fmt.Printf("不执行goto啊")
		return
	end:
		fmt.Printf("goto 到这里啦")
	}
}

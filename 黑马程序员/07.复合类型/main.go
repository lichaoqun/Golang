/*
随机数函数
	函数中的种子只需要设置一次
	如果种子参数一样, 每次程序运行产生随机数都是相同的.
**/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("==== random ====")
	aboutRandom()
	fmt.Println("==== ptr ====")
	aboutPtr()
	fmt.Println("==== array ====")
	aboutArray()
	fmt.Println("==== slice ====")
	aboutSlice()
	fmt.Println("==== map ====")
	aboutMap()
	fmt.Println("==== struct ====")
	aboutStruct()

}

func aboutRandom() {
	{ // - 随机数 这里因为种子是确定的, 所以程序每次启动得到的随机数都是相同的, 如果不希望每次的随机数都相同, 可以使用时间作为随机数种子.
		rand.Seed(99)
		for i := 0; i < 4; i++ {
			// - 随便产生一个随机数
			fmt.Println("产生的随机数为:", rand.Int())

			// - 产生的随机数限制在某个区间
			fmt.Println("产生的随机数为:", rand.Intn(100))
		}
	}
}

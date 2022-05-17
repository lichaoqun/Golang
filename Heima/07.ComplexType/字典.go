/*
map
	map 只有长度, 没有容量
	map 作为函数参数, 是引用传值
	map 初始化时候, 不能使用重复的key
**/
package main

import (
	"fmt"
)

func aboutMap() {
	{ // - 初始化方式
		var m1 map[string]float32 = map[string]float32{"age": 1, "height": 11, "weight": 111}
		fmt.Println("map is :", m1, ", len is :", len(m1))

		// - 自动类型推导初始化
		m2 := map[string]float32{"age": 2, "height": 22, "weight": 222}
		fmt.Println("map is :", m2, ", len is :", len(m2))

		// - 使用make函数
		m3 := make(map[string]float32, 6)
		fmt.Println("map is :", m3, ", len is :", len(m3))

		// - 元素的访问
		m3["age"] = 3
		m3["height"] = 33
		m3["weight"] = 333
		fmt.Println("map is :", m3, ", len is :", len(m3))

		// - 遍历
		for key, value := range m3 {
			fmt.Println("key is", key, ", value is", value)
		}

		// - 判断某个key是否存在
		key := "age"
		value, isExist := m3[key]
		if isExist {
			fmt.Println("in m3, the", key, "is exist, the value is", value)
		}

		// - key 删除
		delete(m3, key)
		fmt.Println("map is :", m3, ", len is :", len(m3))

		// - map 作为函数参数, 传值是引用传值
		modifyMap(m3)
		fmt.Println("map is :", m3, ", len is :", len(m3))

	}
}

func modifyMap(m3 map[string]float32) {
	delete(m3, "height")
}

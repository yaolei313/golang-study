package base

import "fmt"

func StudyAddress() {
	type T struct{ age int }
	mt := map[string]T{}
	mt["John"] = T{age: 29}
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789}

	// mt["John"].age = 30 // error
	t := mt["John"] // 临时变量
	t.age = 30
	mt["John"] = t // 整体修改

	// ma[1][1] = 123      // error
	a := ma[1] // 临时变量
	a[1] = 123
	ma[1] = a // 整体修改
	fmt.Println(ma[1][1], mt["John"].age) // 123 30”
}

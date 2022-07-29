package base

import (
	"fmt"
	"log"
	"reflect"
)

func StudyArray() {
	// 数组的指针
	ps := &[]string{"Go", "C"}
	log.Println("type:", (*ps)[0], (*ps)[1], reflect.TypeOf(ps), reflect.TypeOf(*ps))

	ps0 := &[]string{"Go", "C"}[0]
	fmt.Println(*ps0) // Go

	ps1 := &[]string{"Go", "C"}[1]
	fmt.Println(*ps1) // C

	studyMap1()
	studyMap2()
}

func studyMap1() {
	// map定义后为nil mpa，不能存放键值对，使用make可以初始化
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	fmt.Println(map1)

	capital, ok := map1["US"]
	if ok {
		fmt.Println("us capital is", capital)
	} else {
		fmt.Println("not found in map")
	}
	map1["Japan"] = "Tokyo"
	map1["India"] = "New delhi"
	fmt.Println(map1)
	delete(map1, "France")
	delete(map1, "abc")
	fmt.Println("after modify map1:", map1)

	for key, value := range map1 {
		fmt.Println("key:", key, ", value:", value)
	}
}

func studyMap2() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	// value的地址始终是相同的
	for index, value := range slice {
		fmt.Printf("%v %x\n", value, &value)
		// 只要这样才可以
		item := value
		myMap[index] = &item
	}
	fmt.Println("=====new map=====")
	printMap(myMap)
}

func printMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}

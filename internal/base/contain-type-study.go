package base

import "fmt"

func TestArray() {
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
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}

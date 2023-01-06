package base

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	X, Y int
}

func TestStruct() {
	v := Vertex{10, 11}
	p := &v
	pc := v
	p.X = 1e9
	// {1000000000 11} {10 11} pc:=v实际为复制struct
	fmt.Printf("%v %v \n", v, pc)

	v2 := Vertex{X: 100}
	v3 := Vertex{}
	p2 := &Vertex{
		X: 66,
		Y: 99,
	}
	fmt.Println(v2, v3, p2)
}

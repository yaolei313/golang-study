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
	p.X = 1e9
	fmt.Println(v)

	v2 := Vertex{X: 100}
	v3 := Vertex{}
	p2 := &Vertex{
		X: 66,
		Y: 99,
	}
	fmt.Println(v2, v3, p2)
}

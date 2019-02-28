package main

import (
	"fmt"
	"reflect"
	"go/types"
	"errors"
	"time"
	"math"
)

type NameAware interface {
	GetName() string
}
type Student struct {
	name string
}

func (s Student) GetName() string {
	return s.name
}

func (s *Student) SetName(str string){
	s.name = str
}

func PrintName(aware NameAware) {
	fmt.Println("name is ", aware.GetName())
}

func Swap2(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func Swap1(a, b *int) {
	*a, *b = *b, *a
}

func PrintType(obj interface{}){
	switch i := obj.(type) {
	case types.Nil:
		fmt.Printf("obj is %T\n", i)
	case int:
		fmt.Println("obj is int")
	case bool, string:
		fmt.Println("obj is bool or string")
	case Student:
		fmt.Println("obj is Student")
	case *Student:
		fmt.Println("obj is Student Pointer")
	case uintptr:
		fmt.Println("obj is point")
	default:
		fmt.Println("unknown type")
	}
}

func Factorial(n uint64)(result uint64) {
	if n>0 {
		return n * Factorial(n-1)
	}
	return 1
}

type Handler func(s string) int

func (h Handler) Add(param1 string,param2 int) int{
	return h(param1) + param2
}

func Divide(a,b int) (int,error) {
	if b == 0 {
		return 0,errors.New("b is zero")
	} else {
		return a/b,nil
	}
}

// defer
func sdefer1() int{
	i:=0
	// 1.defer在return之后被调用，但其参数为实时解析
	// 2.多个defer后定义的先被执行
	// 类似finally吧
	defer fmt.Println("second call",i)
	i++
	defer fmt.Println("first call")
	return i
}

func main() {
	// 数据类型
	// uint8 uint16 uint32 uint64
	// int8 int16 int32 int64
	// bool
	// string
	// float32 float64
	// 变量
	// 声明和初始化分离
	var a1, a2 int
	a1, a2 = 1, 3

	// 声明且初始化，类型靠推断
	var b1, b2 = 23, "hello"

	// 局部变量，可以省略var
	c1, c2 := 23, 65535

	var (
		d int
		e int
	)
	d, e = 15, 25
	fmt.Println(a1, a2, b1, b2, c1, c2, d, e)

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
	map1["Japan"]="Tokyo"
	map1["India"]="New delhi"
	fmt.Println(map1)
	delete(map1,"France")
	delete(map1,"abc")
	fmt.Println("after modify map1:",map1)

	// 数组，值类型，长度固定，下标位置不支持负数
	var array1 [2]int
	array1[0] = 100
	array1[1] = 200
	var array2 = [5]int{1, 2, 3, 4, 5}
	var array3 = [...]int{5, 6, 7, 8, 9}
	fmt.Println(array1, "===", array2, "===", array3)

	// 不设置值的话，默认值为0
	var array4 [2][3]int
	var array5 = [3][3]int{
		{11, 12, 13},
		{14, 15, 16},
	}
	// 只能写一个...
	var array6 = [...][3]int{
		{21, 22, 23},
		{24, 25, 26},
	}
	fmt.Println(array4, "===", array5, "===", array6)

	// 切片var identifier []type
	// slice不存储数据，引用类型，可以看做array的section描述
	var array7 = []int{1, 2, 3, 11, 12, 13}
	var slice1 = array7[1:2]
	fmt.Println(len(array7),slice1, reflect.TypeOf(slice1), cap(slice1), len(slice1))
	slice1[0] = 998
	fmt.Println("slice is a reference",array7, slice1)

	var slice2 []int = make([]int,10)
	fmt.Printf("slice2 content:%v cap=%d len=%d\n",slice2, cap(slice2), len(slice2))
	slice2 = append(slice2, 1)
	slice2 = append(slice2, 2,3,4)

	var slice3 = make([]int,5,10)
	fmt.Printf("slice3 content:%v\n", slice3)

	// 类型转换
	var float1 = float32(b1)
	fmt.Println("after convert", float1, reflect.TypeOf(float1))
	var float2 = 5/3
	var float3 = float32(5)/float32(3)
	fmt.Println("除法:",float2,float3)

	// 常量
	const f1 string = "123"
	var f2, f3 int = 10, -2147483648
	fmt.Println(f1, f2, f3)
	// >> 在java及go中都是带符号右移
	fmt.Println(f2>>1, f3>>1)
	f2++

	// 指针，没有分配时值为nil
	var ptr *int
	fmt.Printf("ptr is %v %v\n", ptr, ptr == nil)
	ptr = &f2
	fmt.Printf("*ptr is %d\n", *ptr)
	fmt.Printf("ptr is %v %x\n", ptr, ptr) // v格式带0x

	// 指针的数组
	var ptr2 [6]*int
	for i := 0; i < 6; i++ {
		ptr2[i] = &array7[i]
	}
	fmt.Println("[6] *int ", ptr2)

	// 指针的指针
	var ptr3 **int
	ptr3 = &ptr
	fmt.Println("point:", ptr, ptr3, *ptr3, **ptr3)

	// 条件语句
	if f2 < 100 {
		fmt.Printf("f2 is lower than 100")
	}

	if f2 < 50 {
		fmt.Printf("f2 is lower than 50\n")
	} else {
		fmt.Printf("f2 is greater than or equal to 50")
	}

	switch f2 {
	case 50:
		fmt.Println("switch1 equal 50")
	case 100:
		fmt.Println("switch1 equal 100")
	default:
		fmt.Println("switch1 other value")
	}

	switch {
	case f2 <= 50:
		fmt.Println("switch2 f2 <= 50")
	case f2 <= 100, f2 == 200:
		fmt.Println("switch2 f2 <= 100 or f2 == 200")
	default:
		fmt.Println("switch2 f2 > 100 and f2 != 200")
	}

	student1 := Student{"li bai"}
	student2 := Student{name:"du fu"}
	fmt.Println("=============================")
	PrintName(&student1)
	student1.SetName("li bai2")
	PrintName(&student1)
	PrintType(student2)
	PrintType(&student2)
	fmt.Println("acesss struct field:", student1.name)

	var g1, g2, g3 chan int
	var h1, h2 int
	select {
	case h1 = <-g1:
		fmt.Println("channel g1 received ", h1)
	case g2 <- h2:
		fmt.Println("send to channel g2", h2)
	case h3, ok := (<-g3):
		if ok {
			fmt.Println("channel g3 received ", h3)
		} else {
			fmt.Println("channel g3 is closed")
		}
	default:
		fmt.Println("no communication")
	}

	// 循环语句
	/*for init;condition;post {

	}*/

	// 等同while condition
	/*for condition {

	}*/

	// 等同while true
	/*for {

	}*/

	// for循环的range可以可以遍历slice，map，数组，字符串
	for idx,val := range array7 {
		fmt.Println("idx:",idx,"val:",val)
	}

	for key, value := range map1 {
		fmt.Println("key:", key, ", value:", value)
	}

	// 变量字符串
	for key, value := range b2 {
		fmt.Println("key:", key, ", value:", value)
	}

	fmt.Println("before swap c1:", c1, "c2:", c2)
	Swap1(&c1, &c2)
	fmt.Println("after swap c1:", c1, "c2:", c2)
	Swap2(&c1, &c2)
	fmt.Println("after swap c1:", c1, "c2:", c2)
	c1, c2 = c2, c1
	fmt.Println("after swap c1:", c1, "c2:", c2)

	// for循环的a:=0，与外层的a不是同一个变量，为 for 循环中的局部变量，因此在执行完for循环后，输出a的值仍然为0
	var a int = 0
	fmt.Println("for start")
	for a := 0; a < 10; a++ {
		fmt.Print(a, "\t")
	}
	fmt.Println()
	fmt.Println("for end")
	fmt.Println(a)

	var len_handler Handler = func(s string) int {
		return len(s)
	}
	Handler(len_handler).Add("abc",10);

	// errors.New("")

	// defer学习
	var r1 int = sdefer1()
	fmt.Println("return value", r1)

	// go学习
	go sayHello("hello li bai")
	sayHello("hello world")

	// channel学习
	var chan1 chan int
	fmt.Println("chan1", chan1)
	chan1 = make(chan int)
	defer func(param chan int) { close(param)}(chan1)
	fmt.Println("chan1", chan1)

	// 如果通道不带缓冲，则发送发会阻塞，若带缓冲，则发送阻塞知道被拷贝到缓冲区，若缓冲已满，同样阻塞
	var chan2 = make(chan int,10)
	chan2 <- 1
	chan2 <- 2
	fmt.Println(<-chan2,<-chan2)

	fmt.Println("math:",-math.Sqrt2)
}

func sayHello(s string){
	for i:=0;i<5;i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}





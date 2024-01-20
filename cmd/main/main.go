package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	list := []int{3, 2, 6, 5, 0, 3}
	newList := list[6:]
	fmt.Println(newList)
}

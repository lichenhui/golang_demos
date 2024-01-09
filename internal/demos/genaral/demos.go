package genaral

import (
	"fmt"
	"reflect"
)

func fmtLibs() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 绑方法
func (u User) Hello() {
	fmt.Println("Hello")
}

func reflect12(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s : %v", field.Name, field.Type)
		val := v.Field(i).Interface()
		fmt.Println("value: ", val)
	}
	fmt.Println("---------print functions-------")

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}

// https://leetcode.com/problems/jump-game/description/?envType=study-plan-v2&envId=top-interview-150
func canJump(nums []int) bool {
	firstIndex := 0
	return jumpForward(firstIndex, nums)
}

func jumpForward(currentIndex int, nums []int) bool {
	lastIndex := len(nums) - 1
	if currentIndex == lastIndex {
		return true
	} else if currentIndex < lastIndex {
		currentNode := nums[currentIndex]
		if currentNode == 0 {
			return false
		}
		return jumpForward(currentIndex+currentNode, nums)
	} else {
		// out of index
		return false
	}
}

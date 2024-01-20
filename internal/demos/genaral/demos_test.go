package genaral

import (
	"fmt"
	"testing"
)

func TestFmtLibs(t *testing.T) {
	fmtLibs()
}

func TestReflect(t *testing.T) {
	obj := &User{1, "Li", 20}
	reflect12(*obj)
}

func TestCanJump(t *testing.T) {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func TestMaxProfit(t *testing.T) {

}

func TestMaxProfit188(t *testing.T) {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit188(k, prices))
}

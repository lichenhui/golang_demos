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

func TestWordBreak(t *testing.T) {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}

func TestMinSubArrayLen(t *testing.T) {
	//target := 5
	list := []int{1, 2, 3}
	//fmt.Println(minSubArrayLen(target, list))
	fmt.Println(findComb(1, list, 0, []int{}))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func TestCombine(t *testing.T) {
	fmt.Println(combine(4, 2))
}

func TestMyPow(t *testing.T) {
	fmt.Println(myPow(-2, 3))
}

func TestMinhPathSum(t *testing.T) {
	grids := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grids))
}

func TestUser_HelloniquePathsWithObstacles(t *testing.T) {
	grid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(grid))
}

func TestMinimumTotal(t *testing.T) {
	//triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	triangle := [][]int{{-10}}
	fmt.Println(minimumTotal(triangle))
}

func TestZigzagConvert(t *testing.T) {
	s := "PAYPALISHIRING"
	rows := 3
	fmt.Println(ZigzagConvert(s, rows))
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(BuildTree(preorder, inorder))

}

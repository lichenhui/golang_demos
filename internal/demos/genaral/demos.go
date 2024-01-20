package genaral

import (
	"fmt"
	"reflect"
	"strings"
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

// 123. Best Time to Buy and Sell Stock III
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/?envType=study-plan-v2&envId=top-interview-150
// list all combinations, and compare the profit
func maxProfit(prices []int) int {
	// max profit if have only one transaction
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	// profit if conduct 2 transactions
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit1 := prices[j] - prices[i] // profit in 1st transaction
			// 2nd transaction
			for x := j + 1; x < len(prices); x++ {
				for y := x + 1; y < len(prices); y++ {
					profit2 := prices[y] - prices[x]
					if (profit2 + profit1) > maxProfit {
						maxProfit = profit2 + profit1
					}
				}
			}
		}
	}
	return maxProfit
}

// 188. Best Time to Buy and Sell Stock IV
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/?envType=study-plan-v2&envId=top-interview-150
func maxProfit188(k int, prices []int) int {

	currentProfit := 0
	totalProfitList := make([]int, 0)
	totalProfitList = recursionProfit(k, prices, currentProfit, totalProfitList)
	maxProfit := 0
	for _, profit := range totalProfitList {
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func recursionProfit(k int, prices []int, currentProfit int, totalProfitList []int) []int {
	if k == 0 || len(prices) < 2 {
		return append(totalProfitList, currentProfit)
	}

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			newProfit := prices[j] - prices[i] + currentProfit
			newList := prices[(j + 1):]
			fmt.Printf("buy: %d - sell: %d", prices[i], prices[j])
			fmt.Println("")
			totalProfitList = append(totalProfitList, recursionProfit(k-1, newList, newProfit, totalProfitList)...)
		}
	}
	return totalProfitList
}

// 97. Interleaving String
// https://leetcode.com/problems/interleaving-string/description/?envType=study-plan-v2&envId=top-interview-150
func isInterleave(s1 string, s2 string, s3 string) bool {
	return false
}

// 139. Word Break
// https://leetcode.com/problems/word-break/description/?envType=study-plan-v2&envId=top-interview-150
func wordBreak(s string, wordDict []string) bool {
	result := make([]bool, 0)
	result = workBreakInner(s, wordDict, result)
	for _, v := range result {
		if v {
			return true
		}
	}
	return false
}

func workBreakInner(s string, wordDict []string, result []bool) []bool {
	if len(s) == 0 {
		return append(result, true)
	}
	for _, wd := range wordDict {
		if strings.HasPrefix(s, wd) {
			newS := s[len(wd):]
			return append(result, workBreakInner(newS, wordDict, result)...)
		}
	}
	return append(result, false)
}

// 209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/description/?envType=study-plan-v2&envId=top-interview-150
func minSubArrayLen(target int, nums []int) int {
	// list all the combinations that consists of N elements, 1 <N <= length of number list
	for i := 1; i <= len(nums); i++ {
		currentSum := 0
		sumList := make([]int, 0)
		combs := findComb(i, nums, currentSum, sumList)
		fmt.Println(combs)
		for _, sum := range combs {
			if sum >= target {
				return i
			}
		}
	}
	return 0
}

// list all combinations and calculate their sums
func findComb(n int, nums []int, currentSum int, sumList []int) []int {
	if n == 0 || len(nums) == 0 {
		return append(sumList, currentSum)
	}
	for i := 0; i < len(nums); i++ {
		newSum := currentSum + nums[i]
		newNums := nums[(i + 1):]
		sumList = append(sumList, findComb(n-1, newNums, newSum, sumList)...)
	}
	return sumList
}

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-interview-150
func lengthOfLongestSubstring(s string) int {
	// list all combinations that have target numbers of characters
	for i := len(s); i > 0; i-- {
		combs := make([]string, 0)
		combs = findCombs(i, s, combs)
		for _, str := range combs {
			if !haveDuplicated(str) {
				return len(str)
			}
		}
	}
	return 0
}

func findCombs(length int, s string, combStr []string) []string {
	if len(s) == length {
		return append(combStr, s)
	}
	if len(s) < length {
		return combStr
	}
	subStr := s[0:length]
	combStr = append(combStr, subStr)
	newSubStr := s[length:]
	return append(combStr, findCombs(length, newSubStr, combStr)...)

}

func haveDuplicated(s string) bool {
	if len(s) == 1 {
		return false
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}

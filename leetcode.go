package main

import "fmt"

/*

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。


输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。




*/

func main() {
	nums := []int{1, 1, 2}

	removeDuplicates(nums)

}

func removeDuplicates(nums []int) int {

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {

			if j > i {
				if nums[i] == nums[j] {
					nums = append(nums[:j], nums[j+1:]...)
				}
			} else {
				println(nums[i], nums[j])

			}
		}
	}

	fmt.Println(nums)
	return len(nums)
}

func in(x int, nums []int) bool {
	for _, n := range nums {
		if x == n {
			return true
		}
	}
	return false
}

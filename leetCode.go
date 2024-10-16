package leetCode

//T: O(n^2), S: O(1) 
func TwoSum(nums []int, target int) [2]int {
	
	for i, v := range nums {
		for j := i+1; j < len(nums); j++ {
			if target == v + nums[j] {
				return [2]int{i,j}
			}
		}
	}

	return [2]int{}
}

//T: O(n), S: O(n) 
func TwoSum2(nums []int, target int) [2]int {

	mp := make(map[int]int)
	for ix, num := range nums {
		if ixMapped, isExists := mp[target - num]; isExists {
			return [2]int{ixMapped, ix}
		}
		mp[num] = ix
	}
	return [2]int{}
}

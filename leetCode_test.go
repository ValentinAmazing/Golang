package leetCode

import "testing"

func TestTwoSum(t *testing.T) {
	testTable := []struct {
		nums []int
		target int
		expected [2]int
	}{
		{ //0
			nums: []int{2,7,11,15},
			target: 9,
			expected: [2]int{0,1},
		},
		{ //1
			nums: []int{3, 2, 4},
			target: 6,
			expected: [2]int{1,2},
		},
		{ //2
			nums: []int{3, 3},
			target: 6,
			expected: [2]int{0,1},
		},
		{ //3
			nums: []int{3, 5},
			target: 7,
			expected: [2]int{},
		},
	}

	for i, testCase := range testTable {
		result := TwoSum(testCase.nums, testCase.target)

		if result != testCase.expected {
			t.Errorf( "TwoSum(): Incorrect result. Test case %d, got %v,expected %v", i, result, testCase.expected)
		}


		result = TwoSum2(testCase.nums, testCase.target)

		if result != testCase.expected {
			t.Errorf( "TwoSum2(): Incorrect result. Test case %d, got %v,expected %v", i, result, testCase.expected)
		}
	}
}


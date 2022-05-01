package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
	var ln int = len(nums)
	var i int = 1

	for i < ln {
		if nums[i-1] != nums[i] {
			i++
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			ln--
		}
	}
	return ln
}

func TestRemoveDuplicates(t *testing.T) {
	input1 := []int{1, 1, 2}
	output1 := removeDuplicates(input1)
	assert.Equal(t, output1, 2)

	input2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output2 := removeDuplicates(input2)
	assert.Equal(t, output2, 5)
}

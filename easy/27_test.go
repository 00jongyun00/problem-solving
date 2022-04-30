package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func TestRemoveElement(t *testing.T) {
	input1 := []int{3, 2, 2, 3}
	output1 := removeElement(input1, 3)
	assert.Equal(t, output1, 2)

	input2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	output2 := removeElement(input2, 2)
	assert.Equal(t, output2, 5)
}

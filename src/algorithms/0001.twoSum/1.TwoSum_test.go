package _001_twoSum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	targets := 9
	t.Log(twoSum(nums, targets))
}

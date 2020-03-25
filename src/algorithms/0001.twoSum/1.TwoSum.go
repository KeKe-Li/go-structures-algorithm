package _001_twoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		// 通过查询map，获取a = target - v的序列号
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		// 把v和i的值，存入map
		m[v] = i
	}
	return nil
}

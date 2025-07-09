package prefix_sum

func preProcessSum(nums []int) []int {
	prep := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			prep[i] = v
			continue
		}

		prep[i] = prep[i-1] + v
	}

	return prep
}

func SumElem(nums []int, i, j int) (result int) {
	prep := preProcessSum(nums)

	result = prep[j] - prep[i-1]
	return
}

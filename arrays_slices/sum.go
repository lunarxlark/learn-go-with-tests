package arrays_slices

func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(num2sum ...[]int) []int {
	sums := []int{}
	//sums := make([]int, len(num2sum))
	for _, nums := range num2sum {
		//sums[i] = Sum(nums)
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(num2sum ...[]int) []int {
	sums := []int{}
	for _, nums := range num2sum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}
	}
	return sums
}

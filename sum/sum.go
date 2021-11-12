package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) (sums []int){
	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}
	return
}

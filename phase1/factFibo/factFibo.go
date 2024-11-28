package factFibo

func CalFactorials(num int) int {
	count := 1
	for i := 2; i <= num; i++ {
		count *= i
	}
	return count
}

func CalFibonacci(num int) int {
	if num == 1 {
		return 0
	} else if num == 2 || num == 3 {
		return 1
	}
	old := 1
	new := 1
	for i := 1; i <= num-3; i++ {
		temp := new
		new += old
		old = temp
	}
	return new
}

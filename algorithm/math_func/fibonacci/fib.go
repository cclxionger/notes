package fibonacci

func Fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)

}
func Fib2(n int) int {
	arr := make([]int, n)
	if n == 1 || n == 0 {
		return 1
	}
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < n; i++ {

		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n-1]
}

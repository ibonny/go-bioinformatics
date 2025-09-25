package main

import "fmt"

func fib(memoize map[int]uint64, n int) uint64 {
	k := 2

	if v, ok := memoize[n]; ok {
		return v
	}

	if n == 1 || n == 2 {
		return 1
	}

	memoize[n] = fib(memoize, n-1) + fib(memoize, n-2)*uint64(k)

	return memoize[n]
}

func main() {
	memoize := make(map[int]uint64)

	fmt.Println(fib(memoize, 28))
}

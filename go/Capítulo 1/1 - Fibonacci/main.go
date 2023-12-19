package main

func main() {
	result := fibonacci(50)

	println(result)
}

var cache = map[int]int{0: 0, 1: 1}

func fibonacci(n int) int {
	if _, hasOnCache := cache[n]; !hasOnCache {
		cache[n] = fibonacci(n-1) + fibonacci(n-2)
	}

	return cache[n]
}

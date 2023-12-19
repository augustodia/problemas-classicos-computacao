package main

func main() {
	result := fibonacci(50)

	println(result)
}

var cache = map[int]int{0: 0, 1: 1} // Cache to store the results

func fibonacci(n int) int {
	if _, hasOnCache := cache[n]; !hasOnCache { // If the result is not on cache, calculate it
		cache[n] = fibonacci(n-1) + fibonacci(n-2) // Save the result on cache
	}

	return cache[n] // Return the result
}

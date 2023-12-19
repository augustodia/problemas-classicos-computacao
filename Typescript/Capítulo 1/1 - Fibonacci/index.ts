
const cache = new Map<number, number>([[0, 0], [1, 1]]); // Cache to store the results

function fibonacci(n: number): number {
	if(cache.has(n)) return cache.get(n) as number; // If the result is not on cache, calculate it

	const result = fibonacci(n - 1) + fibonacci(n - 2); 
	cache.set(n, result); // Save the result on cache

	return result; // Return the result
}

const result = fibonacci(50); 

console.log(result)
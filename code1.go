package main

import (
	"fmt"
	"sort"
)

// Binary Search Implementation
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Merge Sort Implementation
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// Fibonacci (Dynamic Programming - Memoization)
var memo = map[int]int{0: 0, 1: 1}

func fibonacci(n int) int {
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

// Graph - BFS Traversal
type Graph struct {
	adjList map[int][]int
}

func (g *Graph) addEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) bfs(start int) []int {
	visited := map[int]bool{}
	queue := []int{start}
	var result []int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if !visited[node] {
			visited[node] = true
			result = append(result, node)
			queue = append(queue, g.adjList[node]...)
		}
	}
	return result
}

// Additional Algorithms and Problems
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// More Sorting Algorithms
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left, right, equal := []int{}, []int{}, []int{}
	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			equal = append(equal, num)
		}
	}
	return append(append(quickSort(left), equal...), quickSort(right)...)
}

// Factorial using recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	arr := []int{10, 3, 2, 5, 7, 6, 1, 4, 8, 9}
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted Array:", sortedArr)

	fmt.Println("Binary Search for 5:", binarySearch(sortedArr, 5))

	fmt.Println("Fibonacci(10):", fibonacci(10))

	graph := Graph{adjList: make(map[int][]int)}
	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 4)
	fmt.Println("BFS Traversal from node 0:", graph.bfs(0))

	fmt.Println("Is 17 Prime?", isPrime(17))
	fmt.Println("GCD of 48 and 18:", gcd(48, 18))
	fmt.Println("LCM of 12 and 15:", lcm(12, 15))
	fmt.Println("Factorial of 5:", factorial(5))
}

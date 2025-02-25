package main

import (
	"container/heap"
	"fmt"
	"math"
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

// Dijkstra's Algorithm
func dijkstra(graph map[int]map[int]int, start int) map[int]int {
	distance := make(map[int]int)
	for node := range graph {
		distance[node] = math.MaxInt32
	}
	distance[start] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item).value
		for neighbor, weight := range graph[current] {
			if newDist := distance[current] + weight; newDist < distance[neighbor] {
				distance[neighbor] = newDist
				heap.Push(pq, &Item{value: neighbor, priority: newDist})
			}
		}
	}
	return distance
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Knapsack Problem (Dynamic Programming)
func knapsack(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

	distance := dijkstra(map[int]map[int]int{
		0: {1: 4, 2: 1},
		1: {3: 1},
		2: {1: 2, 3: 5},
		3: {},
	}, 0)
	fmt.Println("Dijkstra Shortest Path:", distance)

	fmt.Println("Factorial of 5:", factorial(5))
}

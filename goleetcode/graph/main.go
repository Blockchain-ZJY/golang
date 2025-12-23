package main

import "fmt"

func main() {
	graph := make(map[int][]int)
	// 无向图
	addEdge := func(a, b int) {
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	addEdge(0, 1)
	addEdge(0, 2)
	addEdge(1, 2)
	addEdge(2, 3)
	fmt.Println(graph)
	visited := make(map[int]bool)
	dfs(graph, 0, visited)
	bfs(graph, 0)
}
func dfs(graph map[int][]int, start int, visited map[int]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Println("DFS visit:", start)
	for _, neighbor := range graph[start] {
		dfs(graph, neighbor, visited)
	}
}

func bfs(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println("BFS visit:", node)
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

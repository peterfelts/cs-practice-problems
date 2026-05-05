package main

import (
	"fmt"
)

// # Spanning Tree

// Given the adjacency list of an undirected, **connected** graph, `graph`, return a set of edges forming a _spanning tree_.

// A spanning tree is a subset of edges that connects (i.e., "spans") every node and has no cycles.

// Example 1:
// graph = [
//   [1],
//   [0, 2, 5],
//   [1, 3, 4],
//   [2],
//   [2, 5],
//   [1, 4]
// ]
// Output: [[0, 1], [1, 2], [2, 3], [2, 4], [4, 5]]
// There are other valid answers

// Example 2:
// graph = [[1], [0]]
// Output: [[0, 1]]
// A single edge is a valid spanning tree for two nodes.

// Example 3:
// graph = [
//   [1, 2],
//   [0, 2],
//   [0, 1]
// ]
// Output: [[0, 1], [0, 2]]
// There are other valid answers, like [[0, 1], [1, 2]].

// Example 4:
// graph = [[]]
// Output: []
// This graph has a single node and no edges.

// Constraints:

// - `1 ≤ graph.length ≤ 1000`
// - `graph[i].length < 1000`
// - `0 ≤ graph[i][j] < graph.length`
// - The adjacency list is properly formatted, with no parallel edges or self-loops
// - The graph is connected

const UNVISITED int = 0
const VISITED int = 1

func spanningTree() {

	graph := [][]int{
		{1},
		{0, 2, 5},
		{1, 3, 4},
		{2},
		{2, 5},
		{1, 4},
	}

	visited := make([]int, len(graph))
	edges := make(map[[2]int]struct{})

	undirectedGraphDFS(graph, 0, visited, edges)

	out := [][2]int{}
	// finally, convert the map to a slice and return
	for edge, _ := range edges {
		out = append(out, edge)
	}
	fmt.Println(out)
}

func undirectedGraphDFS(graph [][]int, index int, visited []int, edges map[[2]int]struct{}) {

	visited[index] = VISITED
	neighbors := graph[index]

	for _, neighbor := range neighbors {
		if visited[neighbor] == VISITED {
			continue
		} else {
			edges[[2]int{index, neighbor}] = struct{}{}
			undirectedGraphDFS(graph, neighbor, visited, edges)
		}
	}

	fmt.Printf("%d: %d\n", index, neighbors)

}

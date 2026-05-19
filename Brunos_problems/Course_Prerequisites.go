package bruno

import (
	"fmt"
)

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [a, b] indicates that you must take course b first if you want to take course a.

// For example, the pair [0, 1] indicates that to take course 0 you have to first take course 1.

// Return true if you can finish all courses. Otherwise, return false.

// Constraints:

// 1 <= numCourses <= 2000

// 0 <= prerequisites.length <= 5000

// prerequisites[i].length == 2

// All the pairs prerequisites[i] are unique.

// Input/Output Examples
// Example 1: The Happy Path

// Input: numCourses = 2, prerequisites = [[1,0]]

// Output: true

// Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So it is possible.

// Example 2: The Impossible Loop

// Input:

// numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false

// Explanation: To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. This is a deadlock (a cycle).

const DEBUG bool = false

// My approach is to treat this like a graph problem with cycle detection
// build the graph
// walk the graph and detect loops
func canFinish(numCourses int, prerequisites [][]int) bool {

	if len(prerequisites) < 2 {
		return true
	}

	m := make(map[int][]int)

	prereqs := prerequisites

	// build map
	for i := 0; i < len(prereqs); i++ {

		course := prereqs[i][0]
		prereq := prereqs[i][1]

		if _, ok := m[course]; !ok {
			m[course] = []int{prereq}
		} else {
			tmp := m[course]
			tmp = append(tmp, prereq)
			m[course] = tmp
		}
	}

	if DEBUG {
		fmt.Println(m)
	}

	// Walk the graph and find loops
	for course := range m {
		if DEBUG {
			fmt.Printf("root: %d\n", course)
		}
		visiting := make(map[int]struct{})
		visited := make(map[int]struct{})
		if ok, _, _ := dfs(course, m, visiting, visited); !ok {
			return false
		}
	}

	return true
}

func dfs(course int, courseMap map[int][]int, visiting, visited map[int]struct{}) (bool, map[int]struct{}, map[int]struct{}) {

	// if we've already visited this node in the current stack, we've found a cycle
	if _, ok := visiting[course]; ok {
		if DEBUG {
			fmt.Printf("dected VISITING cycle %d. Returning false\n", course)
		}
		return false, visiting, visited
	}
	if DEBUG {
		fmt.Printf("marking %d as VISITING\n", course)
	}
	visiting[course] = struct{}{}

	// get all children of current course
	for _, prereq := range courseMap[course] {
		if DEBUG {
			fmt.Printf("parent: %d, recursing on %d\n", course, prereq)
		}
		ok := false
		ok, visiting, visited = dfs(prereq, courseMap, visiting, visited)
		if !ok {
			return false, visiting, visited
		}
	}

	// mark node as visited if all chirldren returned from DFS
	if _, ok := visited[course]; !ok {
		if DEBUG {
			fmt.Printf("Marking %d as visited\n", course)
		}
		visited[course] = struct{}{}
	}

	// pop the current node from the list of VISITING nodes
	delete(visiting, course)

	return true, visiting, visited
}

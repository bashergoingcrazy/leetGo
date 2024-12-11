package problems


func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _ , p := range prerequisites {
		a, b := p[0], p[1]
		graph[b] = append(graph[b], a)
		inDegree[a]++
	}
	queue := []int{}
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	order := []int{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		order = append(order, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--

			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(order) == numCourses {
		return order
	}
	return []int{}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _ , p := range prerequisites {
		a, b := p[0], p[1]
		graph[b] = append(graph[b], a)
		inDegree[a]++
	}
	queue := []int{}
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	takenCourses := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		takenCourses++

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--

			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return takenCourses == numCourses
}
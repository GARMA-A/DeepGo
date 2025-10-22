package main

func main() {
	graph := NewGrap()

	graph.AddEdge("A", "B")

	graph.AddEdge("A", "C")

	graph.AddEdge("B", "A")

	graph.AddEdge("B", "D")

	graph.AddEdge("D", "F")

	graph.AddEdge("F", "E")
	graph.DFS("A")
	println()
	graph.BFS("A")
}

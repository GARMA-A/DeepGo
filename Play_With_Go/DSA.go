package main

import (
	"fmt"
)

type Graph[K string] struct {
	adjList map[K][]K
}

func NewGrap[K string]() *Graph[K] {
	return &Graph[K]{adjList: make(map[K][]K)}
}

func (g *Graph[K]) AddEdge(src K, data K) {
	newNode := []K{data}
	g.adjList[src] = append(g.adjList[src], newNode...)
}

func (g *Graph[K]) DFS(startNode K) {
	stack := make([]K, 0, 20)
	visited := make(map[K]bool)
	if _, ok := g.adjList[startNode]; !ok {
		panic("start node in the DFS is not exist in the adjList")
	}
	stack = append(stack, startNode)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(top, " ")
		visited[top] = true
		neigbors := g.adjList[top]
		for _, neigbor := range neigbors {
			if !visited[neigbor] {
				stack = append(stack, neigbor)
			}
		}
	}
}

func (g *Graph[K]) BFS(startNode K) {
	deque := make([]K, 0, 20)

	if _, ok := g.adjList[startNode]; !ok {
		panic("start node in the BFS is not exist in the adjList")
	}
	deque = append(deque, startNode)
	visited := make(map[K]bool, 20)
	for len(deque) > 0 {
		top := deque[0]
		deque = deque[1:]

		fmt.Print(top, " ")
		visited[top] = true
		for _, neigbor := range g.adjList[top] {
			if !visited[neigbor] {
				deque = append(deque, neigbor)
			}
		}
	}
}

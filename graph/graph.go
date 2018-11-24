package graph

import (
	"fmt"
	"gekongfei.com/tennessine/gomq/queue"
	"gekongfei.com/tennessine/gomq/stack"
)

type Edge struct {
	source      int
	destination int
	cost        int
	next        *Edge
}

type EdgeList struct {
	head *Edge
}

type Graph struct {
	count      int
	VertexList []*EdgeList
}

func (g *Graph) Init(count int) {
	g.count = count
	g.VertexList = make([]*EdgeList, count)
	for i := 0; i < count; i++ {
		g.VertexList[i] = new(EdgeList)
		g.VertexList[i].head = nil
	}
}

func (g *Graph) AddEdge(source, destination int, cost int) {
	edge := &Edge{
		source:      source,
		destination: destination,
		cost:        cost,
		next:        g.VertexList[source].head,
	}
	g.VertexList[source].head = edge
}

func (g *Graph) AddEdgeUnweighted(source, destination int) {
	g.AddEdge(source, destination, 1)
}

func (g *Graph) AddBiEdge(source, destination int, cost int) {
	g.AddEdge(source, destination, cost)
	g.AddEdge(destination, source, cost)
}

func (g *Graph) AddBiEdgeUnweighted(source, destination int) {
	g.AddBiEdge(source, destination, 1)
}

func (g *Graph) Print() {
	for i := 0; i < g.count; i++ {
		ad := g.VertexList[i].head
		fmt.Print("Vertex ", i, " is connected to : ")
		for ad != nil {
			fmt.Print("[", ad.destination, ad.cost, "]")
			ad = ad.next
		}
		fmt.Println()
	}
}

func (g *Graph) DFSStack() {
	count := g.count
	visited := make([]int, count)
	var curr int
	s := new(stack.Stack)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	visited[1] = 1
	s.Push(1)

	for s.Size() != 0 {
		elem, _ := s.Pop()
		curr = elem.(int)
		fmt.Println(curr)
		head := g.VertexList[curr].head
		for head != nil {
			if visited[head.destination] == 0 {
				visited[head.destination] = 1
				s.Push(head.destination)
			}
			head = head.next
		}
	}
}

func (g *Graph) DFS() {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			visited[i] = 1
			g.DFSRec(i, visited)
		}
	}
}

func (g *Graph) DFSRec(index int, visited []int) {
	head := g.VertexList[index].head
	println(index)
	for head != nil {
		if visited[head.destination] == 0 {
			visited[head.destination] = 1
			g.DFSRec(head.destination, visited)
		}
		head = head.next
	}
}

func (g *Graph) BFS() {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			g.BFSQueue(i, visited)
		}
	}
}

func (g *Graph) BFSQueue(index int, visited []int) {
	var curr int
	q := new(queue.Queue)
	visited[index] = 1
	q.Add(index)

	for q.Size() != 0 {
		elem, _ := q.Remove()
		curr = elem.(int)
		fmt.Println(curr)
		head := g.VertexList[curr].head
		for head != nil {
			if visited[head.destination] == 0 {
				visited[head.destination] = 1
				q.Add(head.destination)
			}
			head = head.next
		}
	}
}

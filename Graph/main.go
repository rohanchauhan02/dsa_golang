package main

import "fmt"

// Graph represents an sdjecency list graph

type Graph struct {
	vertices []*Vertex
}

// Vertex represent a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

// AddEdge

func (g *Graph) AddEdge(from, to int) {
	//get vertex
	//check error
	//add edge
}

//getVertex

//contains

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}

	}
	return false
}

// PrintGraph

func (g *Graph) PrintGraph() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, i := range v.adjacent {
			fmt.Printf("%v: ", i.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddVertex(5)
	test.AddVertex(0)
	test.AddVertex(5)
	test.PrintGraph()
}

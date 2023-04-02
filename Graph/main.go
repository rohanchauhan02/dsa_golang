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
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex ==nil {
		err:= fmt.Errorf("Invalid edge (%v-->%v)",from, to)
		fmt.Println(err.Error())
		return
	}
	if contains(fromVertex.adjacent,to) ||  contains(toVertex.adjacent,from){
		err:= fmt.Errorf("Edge already present (%v-->%v)",from, to)
		fmt.Println(err.Error())
		return
	}
	//add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

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
			fmt.Printf("%v ", i.key)
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
	test.AddEdge(1, 2)
	test.AddEdge(6, 2)
	test.AddEdge(3, 2)
	test.AddEdge(1, 2)
	test.PrintGraph()
}

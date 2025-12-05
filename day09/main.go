package main

import (
	"aoc-2015/common"
	"bytes"
	"container/heap"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Edge struct {
	Destination string
	Weight      int
}

type EdgeHeap []Edge

func (e EdgeHeap) Len() int { return len(e) }

func (e EdgeHeap) Less(i, j int) bool { return e[i].Weight < e[j].Weight }

func (e EdgeHeap) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func (e *EdgeHeap) Push(x any) {
	*e = append(*e, x.(Edge))
}

func (e *EdgeHeap) Pop() any {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func buildAdjacencyMatrix(data []byte) (map[string][]Edge, error) {
	adj := make(map[string][]Edge)

	// trim whitespace
	trimmedNewLine, _ := bytes.CutSuffix(data, []byte{10})
	directions := bytes.SplitSeq(trimmedNewLine, []byte{10})
	for d := range directions {
		strs := strings.Split(string(d), " ")

		src := strs[0]
		dst := strs[2]
		cost, err := strconv.Atoi(strs[len(strs)-1])
		if err != nil {
			return adj, err
		}

		adj[src] = append(adj[src], Edge{
			Destination: dst,
			Weight:      cost,
		})

		adj[dst] = append(adj[dst], Edge{
			Destination: src,
			Weight:      cost,
		})

	}

	return adj, nil

}

func solveMst(adj map[string][]Edge, start string) int {

	var res int
	visited := make(map[string]bool)

	edgeHeap := &EdgeHeap{Edge{
		Destination: start,
		Weight:      0,
	}}

	heap.Init(edgeHeap)

	var order []string

	for edgeHeap.Len() > 0 {

		e := heap.Pop(edgeHeap).(Edge)
		if visited[e.Destination] {
			continue
		}

		fmt.Println(e.Destination, e.Weight)
		res += e.Weight
		visited[e.Destination] = true
		order = append(order, e.Destination)
		for _, dst := range adj[e.Destination] {

			if !visited[dst.Destination] {
				heap.Push(edgeHeap, dst)
			}

		}

		fmt.Println("---")

	}

	fmt.Println(order)

	return res

}

func solveDijkstra(adj map[string][]Edge, start string) int {
	var res int
	visited := make(map[string]bool)

	edgeHeap := &EdgeHeap{Edge{
		Destination: start,
		Weight:      0,
	}}

	heap.Init(edgeHeap)

	for edgeHeap.Len() > 0 {

		e := heap.Pop(edgeHeap).(Edge)
		if visited[e.Destination] {
			continue
		}

		res += e.Weight
		visited[e.Destination] = true
		for _, dst := range adj[e.Destination] {

			if !visited[dst.Destination] {
				dstAddedWt := Edge{
					Destination: dst.Destination,
					Weight:      dst.Weight + e.Weight,
				}
				heap.Push(edgeHeap, dstAddedWt)
			}
		}
	}

	return res

}

func main() {

	// filePath := "./inputExample.txt"
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	adj, err := buildAdjacencyMatrix(data)
	if err != nil {
		log.Fatal(err)
	}

	var startingCity string
	for k := range adj {
		startingCity = k
		break
	}

	res := solveMst(adj, startingCity)
	fmt.Println(res)

	res2 := math.MaxInt
	for k := range adj {
		v := solveDijkstra(adj, k)
		fmt.Println(v)
		res2 = min(res2, v)
	}

	fmt.Println(res2)

}

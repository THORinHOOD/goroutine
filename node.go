package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Node struct {
	links []*Node
	value int
}

func (node Node) String() string {
	a := func(node *Node) string {
		if node == nil {
			return "{}"
		}
		return node.String()
	}
	strs := make([]string, 0)
	for _, link := range node.links {
		strs = append(strs, a(link))
	}

	return fmt.Sprintf("{%s}[%d]", strings.Join(strs[:], ","), node.value)
}

func randomBinTree(depth int, width int) *Node {
	if depth == 0 || rand.Intn(20) > 18 {
		return nil
	}

	node := Node{make([]*Node, 0), rand.Intn(10)}
	for i := 0; i < width; i++ {
		node.links = append(node.links, randomBinTree(depth-1, width))
	}
	return &node
}

func (node *Node) Sum() (int, time.Duration) {
	start := time.Now()
	if node == nil {
		return 0, time.Since(start)
	}

	sum := node.value
	for _, link := range node.links {
		val, _ := link.Sum()
		sum += val
	}
	return sum, time.Since(start)
}

func (node *Node) SumParallel() (int, time.Duration) {
	start := time.Now()
	if node == nil {
		return 0, time.Since(start)
	}

	sum := node.value
	var wg sync.WaitGroup
	summation := func(node *Node) {
		defer wg.Done()
		val, _ := node.Sum()
		sum += val
	}
	for _, link := range node.links {
		wg.Add(1)
		go summation(link)
	}
	wg.Wait()
	return sum, time.Since(start)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	node := randomBinTree(5, 25)
	fmt.Println("Start to calc...")
	fmt.Println(node.Sum())
	fmt.Println(node.SumParallel())
}

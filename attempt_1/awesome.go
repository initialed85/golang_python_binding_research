package main

import (
	"C"
	"sort"
	"fmt"
	"sync"
	"math"
)

var count int
var mtx sync.Mutex

func Add(a, b int) int { return a + b }

func Cosine(x float64) float64 { return math.Cos(x) }

func Sort(vals []int) { sort.Ints(vals) }

func Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

func main() {}

package connectivity

import (
	"math/rand"
	"time"
)

// monte carlo simulator for percolation
// once initiated, the simulator fills up
// blocks randomly and check for percolation
// on each step.
// the simulation stops when a percolation is found
// and the simulation will return the used steps that
// produced the percolation
type Simulator interface {

	// run simulation til a percolation is found
	// return the used steps to produce a percolation
	Simulate() int64

	// mark the index n as open. search for neighbours,
	// union with those who are also open.
	mark(n int)

	// clear states of the last simulation, init all values.
	// size of the grid and side length stay the same
	Clear()
}

// percolation simulator
type PercolationSimulator struct {
	Size    int
	Marked  []bool
	l       int
	connect *WeightedCompressed
}

// ctor for percolation simulator
// eats a integer representing the side length of the
// square. e.g. N -> NxN sqaure
// the square is initially completely black or blocked.
func NewPercolationSimulator(N int) *PercolationSimulator {
	result := &PercolationSimulator{
		Size:    N * N,
		Marked:  make([]bool, N*N),
		l:       N,
		connect: NewWeightedCompression(N*N + 2),
	}
	for i := 0; i != N; i++ {
		result.connect.Union(i, result.Size)
	}
	for i := N * (N - 1); i != N*N; i++ {
		result.connect.Union(i, result.Size+1)
	}
	return result
}

// every call to the simulate will generate a random order of
// how the blocks will be marked, and they will be marked according
// to that order utill we found the percolation
// returns the steps taken to find the percolation
func (p *PercolationSimulator) Simulate() int64 {

	p.Clear() //clear states from last simulation

	perm := getPermutation(p.Size)
	steps := 0

	for _, idx := range perm {
		p.mark(idx)
		steps++
		if p.connect.Find(p.Size, p.Size+1) {
			break
		}
	}

	return int64(steps)
}

// return a slice of pseudo-random permutation of [0,n)
func getPermutation(n int) []int {
	seed := time.Now().UnixNano() % 271828182833
	r := rand.New(rand.NewSource(seed))
	return r.Perm(n)
}

// mark (paint) the block as white
func (p *PercolationSimulator) mark(n int) {
	if p.Marked[n] {
		return
	}
	p.Marked[n] = true
	neighbors := getDirectNeighbours(n, p.l)
	for _, adj := range neighbors {
		if p.Marked[adj] {
			p.connect.Union(n, adj)
		}
	}
}

// direct neighbours of a pixel
func getDirectNeighbours(n, l int) []int {
	var result []int
	if (n / l) < 1 {
		if n%l == 0 {
			result = append(result, n+1)
			result = append(result, n+l)
			return result
		}
		if n%l == l-1 {
			result = append(result, n-1)
			result = append(result, n+l)
			return result
		}
		result = append(result, n-1)
		result = append(result, n+1)
		result = append(result, n+l)
		return result
	}
	if (n / l) >= l-1 {
		if n%l == 0 {
			result = append(result, n-l)
			result = append(result, n+1)
			return result
		}
		if n%l == l-1 {
			result = append(result, n-l)
			result = append(result, n-1)
			return result
		}
		result = append(result, n-l)
		result = append(result, n-1)
		result = append(result, n+1)
		return result
	}
	if n%l == 0 {
		result = append(result, n-l)
		result = append(result, n+1)
		result = append(result, n+l)
		return result
	}
	if n%l == l-1 {
		result = append(result, n-l)
		result = append(result, n-1)
		result = append(result, n+l)
		return result
	}
	result = append(result, n-l)
	result = append(result, n-1)
	result = append(result, n+1)
	result = append(result, n+l)
	return result
}

// clear all states
func (p *PercolationSimulator) Clear() {
	for i := range p.Marked {
		p.Marked[i] = false
	}
	p.connect = NewWeightedCompression(p.Size + 2)
	for i := 0; i != p.l; i++ {
		p.connect.Union(i, p.Size)
	}
	for i := p.l * (p.l - 1); i != p.Size; i++ {
		p.connect.Union(i, p.Size+1)
	}
}

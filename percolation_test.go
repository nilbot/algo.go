package connectivity

import (
	"log"
	"runtime"
	"testing"
)

func TestNewSimulatorSize(t *testing.T) {
	s := NewPercolationSimulator(5)
	if s.Size != 25 {
		t.Errorf("Expected 25, got %v", s.Size)
	}
}

func TestRunMonteCarlo(t *testing.T) {

	CPUS := runtime.NumCPU()

	runtime.GOMAXPROCS(CPUS)

	n := 1000000

	workload := n / runtime.NumCPU()

	results := make(chan int64, CPUS)
	for c := 0; c < CPUS; c++ {
		go func() {
			s := NewPercolationSimulator(5)
			var value int64
			for i := 0; i < workload; i++ {
				value += s.Simulate() // sum of steps
			}
			results <- int64(value)
			log.Printf("CPU %v returned steps %v out of workload %v", c, value, 25*workload)
		}()
	}
	var total int64
	for i := 0; i < CPUS; i++ {
		total += <-results
	}
	log.Printf("ran %v simulations, got result %v", n, float64(total)/float64(25*n))
}

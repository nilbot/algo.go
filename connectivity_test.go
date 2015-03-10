package connectivity

import (
	"bufio"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

// download test data from internet
const (
	tiny   string = "http://algs4.cs.princeton.edu/15uf/tinyUF.txt"
	medium string = "http://algs4.cs.princeton.edu/15uf/mediumUF.txt"
	huge   string = "http://algs4.cs.princeton.edu/15uf/largeUF.txt"
)

func loadTestDataFromWeb() []TestData {

	c := make(chan TestData)
	go func() { c <- parse(tiny) }()
	go func() { c <- parse(medium) }()
	go func() { c <- parse(huge) }()

	var results []TestData
	timeout := time.After(1 * time.Second)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			log.Println("timed out fetching data")
			return results
		}

	}
	return results
}

func parse(url string) (result TestData) {

	cl := &http.Client{}
	r, err := cl.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Body.Close()

	s := bufio.NewScanner(r.Body)
	s.Split(bufio.ScanWords)
	var pairs []Pair
	s.Scan()
	g, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Panicf("first generator parse error %v", g, err)
	}
	for s.Scan() {
		var pair Pair
		pair.Left, err = strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("parse left %v failed", pair.Left, err)
		}
		s.Scan()
		pair.Right, err = strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("parse right %v failed", pair.Right, err)
		}
		pairs = append(pairs, pair)
	}

	result.Generator = g
	result.Pairs = pairs
	return result
}

type Pair struct {
	Left  int
	Right int
}

type TestData struct {
	Generator int
	Pairs     []Pair
}

func TestJustTest(t *testing.T) {
	loadStart := time.Now()
	n := loadTestDataFromWeb()
	log.Printf("loading used %v", time.Since(loadStart))
	for _, content := range n {
		log.Printf("Generator %v has %v pairs", content.Generator, len(content.Pairs))
		algoStart := time.Now()
		test := NewWeightedCompression(content.Generator)
		for idx, p := range content.Pairs {
			// log.Printf("%vth pair: %v and %v", idx, p.Left, p.Right)
			test.Union(p.Left, p.Right)
			if !test.Find(p.Left, p.Right) {
				t.Errorf("%v and %v are expected to be in the same component, this is the %vth union operation", p.Left, p.Right, idx)
			}
		}
		log.Printf("Weighted Compressed UnionFind used %v", time.Since(algoStart))
	}
}

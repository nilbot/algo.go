package connectivity

//////////////////////////////////////////////////////////
//  Improved UnionFind with weight and compression
//////////////////////////////////////////////////////////

// WeightedCompressed holds a parent array to track parent of each note, and a
// size array for weights in terms of number of notes in that component.
type WeightedCompressed struct {
	parent []int
	size   []int
}

// NewWeightedCompression initialize an empty WeightedCompressed struct
func NewWeightedCompression(n int) *WeightedCompressed {
	p := make([]int, n)
	sz := make([]int, n)
	for i := 0; i != n; i++ {
		p[i] = i
		sz[i] = 1
	}
	return &WeightedCompressed{
		p,
		sz,
	}
}

// UnionFind interface defines 2 public operations for the algorithm
type UnionFind interface {
	// Find also known as connected -> bool in literature, check if the 2 notes
	// are in the same component
	Find(a, b int) bool
	// Union connect the 2 notes, make them belong to the same component
	Union(a, b int)
}

// Find delivers the check by comparing roots
func (w *WeightedCompressed) Find(a, b int) bool {
	return w.root(a) == w.root(b)
}

func (w *WeightedCompressed) root(pos int) int {
	for pos != w.parent[pos] {
		// path compression
		w.parent[pos] = w.parent[w.parent[pos]]
		pos = w.parent[pos]
	}
	return pos
}

// Union applies weighted union by balancing based on the size of components
func (w *WeightedCompressed) Union(a, b int) {
	i := w.root(a)
	j := w.root(b)
	if i == j {
		return
	}
	if w.size[i] < w.size[j] {
		w.parent[i] = j
		w.size[j] += w.size[i]
	} else {
		w.parent[j] = i
		w.size[i] += w.size[j]
	}
}

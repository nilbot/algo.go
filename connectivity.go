package connectivity

//////////////////////////////////////////////////////////
//  Improved UnionFind with weight and compression
//////////////////////////////////////////////////////////

// in order to calculate weight of the tree we maintain
// a second array storing the size of the tree at idx
type WeightedCompressed struct {
	parent []int
	size   []int
}

// constructor for weighted compressed connectivity
func NewWeightedCompression(n int) *WeightedCompressed {
	p := make([]int, n)
	sz := make([]int, n)
	for i := 0; i != n; i++ {
		p[i] = i
		sz[i] = i
	}
	return &WeightedCompressed{
		p,
		sz,
	}
}

// finally we hammer the final nail on the API
type UnionFind interface {
	Find(a, b int) bool
	Union(a, b int)
}

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

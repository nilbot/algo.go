package connectivity

// naive findQuery, O(n) in single component
func (s *Component) FindQueryInSingleSetNaive(a, b int) bool {
	//naive
	var first, second bool
	for _, it := range *s {
		if a == it {
			first = true
		}
		if b == it {
			second = true
		}
	}
	return first && second
}

type Component []int
type Components []Component

// return true if a and b are in the same component
func (s *Components) FindQuery(a, b int) bool {
	for _, c := range *s {
		if c.FindQueryInSingleSetNaive(a, b) {
			return true
		}
	}
	return false
}

// return queried Component in which the q resides
// and the index of q in the component
func (s *Components) queryComponent(q int) (int, Component) {
	for i, c := range *s {
		for _, d := range c {
			if d == q {
				return i, c
			}
		}
	}
	return -1, nil
}

type Connectivity struct {
	Size     int
	elements []int
	Components
}

// Construct a array of number 0...N-1
func NewConnectivity(N int) *Connectivity {
	result := &Connectivity{
		Size: N,
	}
	items := make([]int, N)
	components := make([]Component, N)
	for index := 0; index != N; index++ {
		items[index] = index
		components[index] = []int{index}
	}
	result.elements = items
	result.Components = components
	return result
}

// union
// reduce the []component dimension
// and and merge the related 2 components
func (c *Components) Union(a, b int) {
	iA, _ := c.queryComponent(a)
	iB, B := c.queryComponent(b)
	(*c)[iA] = append((*c)[iA], B...)
	copy((*c)[iB:], (*c)[iB+1:])
	(*c)[len(*c)-1] = nil
	*c = (*c)[:len(*c)-1]
}

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

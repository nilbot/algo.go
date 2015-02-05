package connectivity

import (
	"sort"
	"testing"
)

// test set is a struct of set/bag of items, plus
// the test subject -> 2 items
// and expected output -> true or false
type TestSet struct {
	Exp      bool
	SubjectA int
	SubjectB int
	Set      Component
}

// prepare test sets
func getTestSets() *[]TestSet {
	return &[]TestSet{
		{
			Exp:      true,
			SubjectA: 2,
			SubjectB: 3,
			Set:      []int{2, 3, 6, 7},
		},
		{
			Exp:      true,
			SubjectA: 1,
			SubjectB: 4,
			Set:      []int{1, 4, 5},
		},
		{
			Exp:      false,
			SubjectA: 2,
			SubjectB: 9,
			Set:      []int{1, 2, 3, 4, 5, 7},
		},
	}
}

func TestFindQueryInSet(t *testing.T) {
	testSets := getTestSets()
	for _, test := range *testSets {
		exp := test.Exp
		inputA := test.SubjectA
		inputB := test.SubjectB
		result := test.Set.FindQueryInSingleSetNaive(inputA, inputB)
		if result != exp {
			t.Errorf("expected %q, got %q", exp, result)
		}
	}
}

// we have multiple components (sets)
// FindQuery is expected to Find if the 2 input items are
// in the same component.
// We need a struct for Components just like our previous TestSet
type TestConnectivity struct {
	Expected bool
	A        int
	B        int
	Components
}

func getTestComponents() []Component {
	return []Component{
		Component{2, 3, 6, 7}, Component{1, 4}, Component{5, 8, 9},
	}
}

func getTestSuits() *[]TestConnectivity {
	return &[]TestConnectivity{
		{
			Expected:   true,
			A:          2,
			B:          3,
			Components: getTestComponents(),
		},
		{
			Expected:   true,
			A:          1,
			B:          4,
			Components: getTestComponents(),
		},
		{
			Expected:   false,
			A:          2,
			B:          9,
			Components: getTestComponents(),
		},
	}
}

// test findquery in labyrinth
func TestFindQuery(t *testing.T) {
	testSuits := getTestSuits()
	for _, test := range *testSuits {
		exp := test.Expected
		a := test.A
		b := test.B
		result := test.Components.FindQuery(a, b)
		if result != exp {
			t.Errorf("expected %q, got %q", exp, result)
		}
	}
}

// helper func for checking 2 components
// components sorted to deliver the false negative result
// as soon as possible
func checkComponent(a, b Component) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) == len(b) {
		sort.Ints(a)
		sort.Ints(b)
		index := 0
		for index != len(a) {
			if a[index] != b[index] {
				return false
			}
			index++
		}
		return true
	}
	return false
}

// test helper func queryComponent
func TestQueryComponent(t *testing.T) {
	c := Components(getTestComponents())

	if _, got := c.queryComponent(1); !checkComponent(got, c[1]) {
		t.Errorf("expected component %v, got %v", c[1], got)
	}
	if _, got := c.queryComponent(2); !checkComponent(got, c[0]) {
		t.Errorf("expected component %v, got %v", c[0], got)
	}
	if _, got := c.queryComponent(5); !checkComponent(got, c[2]) {
		t.Errorf("expected component %v, got %v", c[2], got)
	}
	// not found digit should return nil instead of empty set, because it would
	// mess with FindQuery
	if _, got := c.queryComponent(0); !checkComponent(got, nil) {
		t.Errorf("expected component %v, got %v", nil, got)
	}
}

// test constructor
func TestConstructorNewConnectivity(t *testing.T) {
	test := NewConnectivity(10)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := 0
	for index != len(test.elements) {
		if expected[index] != test.elements[index] {
			t.Errorf("expected %v, got %v", expected[index], test.elements[index])
		}
		index++
	}
}

// test union
// todo: check if the Union skips cases where two numbers are already connected
func TestUnion(t *testing.T) {
	components := Components(getTestComponents())
	components.Union(1, 5)
	if len(components) != 2 {
		t.Errorf("expected reduced component of length %v, got %v", 2, len(components))
	}

	for _, c := range components {
		if len(c) != 4 {
			if b := checkComponent(c, []int{1, 4, 5, 8, 9}); !b {
				t.Errorf("expected matched component %q, got %q", []int{1, 4, 5, 8, 9}, c)
			}
		}
	}
}

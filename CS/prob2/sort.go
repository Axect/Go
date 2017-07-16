package prob2

import (
	"fmt"
	"math"
	"sort"
)

// =============================================================================
// Basic Type
// =============================================================================

// PhysicsTest includes name & score
type PhysicsTest struct {
	Name  string
	Score int
}

// String() is just string method
func (p PhysicsTest) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Score)
}

// PhysicsList : Let's type to list
type PhysicsList []PhysicsTest

// NewContainer contains rank
type NewContainer struct {
	Inform PhysicsTest
	Rank   int
}

// =============================================================================
// Methods for Physics List
// =============================================================================
func (p PhysicsList) Len() int           { return len(p) }
func (p PhysicsList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PhysicsList) Less(i, j int) bool { return p[i].Score < p[j].Score }

// Mean is mean
func (p PhysicsList) Mean() float64 {
	s := 0
	l := p.Len()
	for _, elem := range p {
		s += elem.Score
	}
	return float64(s) / float64(l)
}

// Std is std
func (p PhysicsList) Std() float64 {
	s := float64(0)
	l := p.Len()
	m := p.Mean()
	for _, elem := range p {
		s += math.Pow(float64(elem.Score)-m, 2)
	}
	return math.Sqrt(s / float64(l))
}

// =============================================================================
// Methods for NewContainer
// =============================================================================
func (Elem NewContainer) String() string {
	return fmt.Sprintf("%s, %d, %d\n", Elem.Inform.Name, Elem.Inform.Score, Elem.Rank)
}

// AssignRank - rank
func (Elem *NewContainer) AssignRank(n int) { Elem.Rank = n }

// AssignInform - inform
func (Elem *NewContainer) AssignInform(value PhysicsTest) { Elem.Inform = value }

// Rank ranks list
func Rank(SortedArray PhysicsList) []NewContainer {
	length := SortedArray.Len()
	Temp := make([]NewContainer, length, length)
	for i := range SortedArray {
		Temp[i].AssignRank(i + 1)
		Temp[i].AssignInform(SortedArray[i])
	}
	return Temp
}

// =============================================================================
// Action
// =============================================================================

// DoSort do sort
func DoSort() {
	fmt.Println("---------------------------------------")
	fmt.Println("Sort List")
	fmt.Println("---------------------------------------")
	A := PhysicsList{
		{"Schrodinger", 64},
		{"Einstein", 90},
		{"Feynman", 98},
		{"Neumann", 100},
		{"Dirac", 92},
		{"Bohr", 90},
		{"Fermi", 91},
		{"Heisenberg", 72},
		{"Pauli", 88},
		{"Newton", 93},
		{"Leibniz", 93},
		{"Planck", 52},
	}
	sort.Sort(sort.Reverse(A))
	//Result := Rank(A)
	fmt.Printf("Average: %f\n", A.Mean())
	fmt.Printf("Std: %f\n ", A.Std())
}

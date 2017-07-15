package prob2

import (
	"fmt"
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

// =============================================================================
// Methods for Physics List
// =============================================================================
func (p PhysicsList) Len() int           { return len(p) }
func (p PhysicsList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PhysicsList) Less(i, j int) bool { return p[i].Score < p[j].Score }

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
	fmt.Println(A)
}

package prob2

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/Axect/Go/Package/csv"
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
	return math.Sqrt(s / (float64(l) - 1)) // -1 means sample
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
		if i > 0 && SortedArray[i].Score == SortedArray[i-1].Score {
			Temp[i].AssignRank(int(Temp[i-1].Rank))
		} else {
			Temp[i].AssignRank(int(i + 1))
		}
		Temp[i].AssignInform(SortedArray[i])
	}
	return Temp
}

// ConvertString converts string list
func ConvertString(C []NewContainer) [][]string {
	length := len(C)
	Temp := make([][]string, length, length)
	for i, elem := range C {
		Temp[i] = []string{fmt.Sprint(elem.Inform.Name), fmt.Sprint(elem.Inform.Score), fmt.Sprint(elem.Rank)}
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

	// Sorting
	sort.Sort(sort.Reverse(A))

	// Statistic
	fmt.Printf("Average: %f\n", A.Mean())
	fmt.Printf("Std: %f\n ", A.Std())

	// Add Rank
	Result := Rank(A)

	// Write CSV
	Data := ConvertString(Result)
	csv.Write(Data, "../Data/physics.csv")

	// Read CSV
	List := csv.Read("../Data/physics.csv")
	Physicsists := make([]string, len(List), len(List))
	Scores, Ranks := make([]int64, len(List), len(List)), make([]int64, len(List), len(List))
	for i, group := range List {
		Physicsists[i] = group[0]
		Scores[i], _ = strconv.ParseInt(group[1], 0, 64)
		Ranks[i], _ = strconv.ParseInt(group[2], 0, 64)
	}
	fmt.Println(Ranks)
}

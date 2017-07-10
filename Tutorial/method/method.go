package method

// Person is sample struct
type Person struct {
	Name    string
	Account int
}

// Add is sample mathod (Pointer Receiver - It can change value)
func (a *Person) Add() {
	for i := 1; i <= 1e+06; i++ {
		a.Account++
	}
}

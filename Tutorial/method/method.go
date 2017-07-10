package method

// Person is sample struct
type Person struct {
	Name    string
	Account int
}

// Add is sample mathod (Pointer Receiver - It can change value)
func (a *Person) Add() {
	a.Account++
}

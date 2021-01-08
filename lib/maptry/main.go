package main

type Person struct {
	Name string
}

func main() {
	pplValue := make(map[string]Person)
	pplPointer := make(map[string]*Person)

	pplValue["1"].Name = "me"
	pplPointer["1"].Name = "her"
}

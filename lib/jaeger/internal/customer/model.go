package customer

type Customer struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

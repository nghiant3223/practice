package route

type Route struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	X    int    `json:"from"`
	Y    int    `json:"to"`
}

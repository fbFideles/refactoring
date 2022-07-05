package item

// New define um construtor
func New(name string, quantity int) *Item {
	return &Item{
		Name:     name,
		Quantity: quantity,
	}
}

type Item struct {
	Name     string
	Quantity int
}

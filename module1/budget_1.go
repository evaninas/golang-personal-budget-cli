package module1

// Budget stores budget information
type Budget struct {
	Max   float32
	Items []Item
}

// Item stores item information
type Item struct {
	Items       []Item
	Description string
	Price       float32
}

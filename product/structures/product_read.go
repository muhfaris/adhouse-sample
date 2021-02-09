package structures

// ProductRead is wrap product id
type ProductRead struct {
	ID   []int  `schema:"id,omitempty" json:"id,omitempty"`
	Name string `schema:"name,omitempty" json:"name,omitempty"`
}

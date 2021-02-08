package domain

type Product struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	QTY  int    `json:"qty,omitempty"`
}

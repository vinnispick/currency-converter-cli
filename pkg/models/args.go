package models

type Args struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	List   bool    `json:"list,omitempty"`
}

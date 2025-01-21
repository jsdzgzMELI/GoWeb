package domain

// TicketAttributes is an struct that represents a ticket
type TicketAttributes struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Country string  `json:"country"`
	Hour    string  `json:"hour"`
	Price   float64 `json:"price"`
}

// Ticket represents a ticket
type Ticket struct {
	Id         int              `json:"id"`
	Attributes TicketAttributes `json:"attributes"`
}

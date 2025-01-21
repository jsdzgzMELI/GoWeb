package repository

import (
	"fmt"

	"github.com/jsdzgzMELI/Desafio-Cierre/internal/domain"
)

// RepositoryTicket represents the repository interface for tickets
type RepositoryTicket interface {
	// GetAll returns all the tickets
	// Get(ctx context.Context) (t map[int]TicketAttributes, err error)
	GetTickets() (map[int]domain.Ticket, error)
	DeleteTicket(int) error
	GetById(int) (domain.Ticket, error)
	AddTicket(t *domain.TicketAttributes) (err error)
	PatchTicket(t domain.TicketAttributes, id int) (err error)
	UpdateTicket(t domain.TicketAttributes, id int) (err error)
	GetTicketsByDestinationCountry(country string) (t map[int]domain.Ticket, err error)
	GetTicketProportion(country string) (float64, error)

	// GetTicketByDestinationCountry returns the tickets filtered by destination country
	// GetTicketByDestinationCountry(ctx context.Context, country string) (t map[int]TicketAttributes, err error)
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type repositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]domain.Ticket

	// lastId represents the last id of the ticket
	lastId int
}

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(dbFile map[int]domain.Ticket, lastId int) RepositoryTicket {
	return &repositoryTicketMap{
		db:     dbFile,
		lastId: lastId,
	}
}

// GetAll returns all the tickets
func (r *repositoryTicketMap) GetAll() (t map[int]domain.Ticket, err error) {
	// create a copy of the map
	// t = make(map[int]domain.Ticket, len(r.db))
	// for k, v := range r.db {
	// 	t[k] = v
	// }

	// return
	t = (*r).db

	return
}

func (r *repositoryTicketMap) GetTickets() (t map[int]domain.Ticket, err error) {
	// create a copy of the map
	t = (*r).db

	return
}

func (r *repositoryTicketMap) GetById(id int) (t domain.Ticket, err error) {
	if t, ok := r.db[id]; ok {
		return t, nil
	}
	return domain.Ticket{}, fmt.Errorf("ticket not found")
}

func (r *repositoryTicketMap) GetLastId() int {
	return r.db[len(r.db)-1].Id
}

func (r *repositoryTicketMap) GetTicketProportion(country string) (proportion float64, err error) {
	count, err := r.GetTicketsByDestinationCountry(country)
	if err != nil {
		return 0, err
	}

	proportion = float64(len(count)) * 100.0 / float64(len(r.db))

	return proportion, nil
}

func (r *repositoryTicketMap) AddTicket(t *domain.TicketAttributes) (err error) {
	// ticket := r.db[r.GetLastId()+1]

	r.lastId++
	var ticket domain.Ticket
	ticket.Attributes = *t
	ticket.Id = r.lastId
	(*r).db[ticket.Id] = ticket
	fmt.Println(r.db)

	return

	// r.lastId++

	// var ticket domain.Ticket
	// ticket.Attributes = *t
	// ticket.Id = r.lastId

	// fmt.Println(r.lastId)

	// r.db[ticket.Id] = ticket

	// return
}

func (r *repositoryTicketMap) DeleteTicket(id int) (err error) {
	if _, ok := r.db[id]; ok {
		delete(r.db, id)
		return nil
	}
	return fmt.Errorf("ticket not found")
}

func (r *repositoryTicketMap) PatchTicket(t domain.TicketAttributes, id int) (err error) {
	if _, ok := r.db[id]; ok {
		if t.Name != "" {
			ticket := r.db[id]
			ticket.Attributes.Name = t.Name
			r.db[id] = ticket
		}
		if t.Email != "" {
			ticket := r.db[id]
			ticket.Attributes.Email = t.Email
			r.db[id] = ticket
		}
		if t.Country != "" {
			ticket := r.db[id]
			ticket.Attributes.Country = t.Country
			r.db[id] = ticket
		}
		if t.Hour != "" {
			ticket := r.db[id]
			ticket.Attributes.Name = t.Name
			r.db[id] = ticket
		}
		if t.Price != 0 {
			ticket := r.db[id]
			ticket.Attributes.Price = t.Price
			r.db[id] = ticket
		}
		return nil
	}
	return fmt.Errorf("Ticket not found")
}

func (r *repositoryTicketMap) UpdateTicket(t domain.TicketAttributes, id int) (err error) {
	if _, ok := r.db[id]; ok {
		err = r.ValueCheck(t)
		if err != nil {
			return err
		}
		ticket := r.db[id]
		ticket.Attributes = t
		r.db[id] = ticket
		return nil
	}
	return fmt.Errorf("Ticket not found")
}

func (r *repositoryTicketMap) ValueCheck(tkt domain.TicketAttributes) (err error) {
	if tkt.Name == "" || tkt.Email == "" || tkt.Country == "" || tkt.Hour == "" || tkt.Price == 0 {
		return fmt.Errorf("There are fields with missing values")
	}
	return nil
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *repositoryTicketMap) GetTicketsByDestinationCountry(country string) (t map[int]domain.Ticket, err error) {
	// create a copy of the map
	t = make(map[int]domain.Ticket)
	for k, v := range r.db {
		if v.Attributes.Country == country {
			t[k] = v
		}
	}
	if len(t) == 0 {
		return nil, fmt.Errorf("No tickets found for the given country")
	}
	return
}

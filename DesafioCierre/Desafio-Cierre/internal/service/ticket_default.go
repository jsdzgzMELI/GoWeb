package service

import (
	"github.com/jsdzgzMELI/Desafio-Cierre/internal/domain"
	"github.com/jsdzgzMELI/Desafio-Cierre/internal/repository"
)

type ServiceTicket interface {
	// GetTotalAmountTickets returns the total amount of tickets
	// GetTotalAmountTickets() (total int, err error)
	GetTickets() (map[int]domain.Ticket, error)
	DeleteTicket(int) error
	GetById(int) (domain.Ticket, error)
	AddTicket(t *domain.TicketAttributes) (err error)
	PatchTicket(t domain.TicketAttributes, id int) (err error)
	UpdateTicket(t domain.TicketAttributes, id int) (err error)
	GetTicketsByDestinationCountry(country string) (t map[int]domain.Ticket, err error)
	GetTicketProportion(country string) (float64, error)
	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	// ...

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	// ...
}

// ServiceTicketDefault represents the default service of the tickets
type serviceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.RepositoryTicket) ServiceTicket {
	return &serviceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *serviceTicketDefault) GetALl() (total int, err error) {
	total = 0

	// get all tickets
	_, err = s.rp.GetTickets()
	if err != nil {
		return
	}

	return
}

func (s *serviceTicketDefault) GetTickets() (t map[int]domain.Ticket, err error) {
	t, err = s.rp.GetTickets()
	if err != nil {
		return
	}
	return
}

func (s *serviceTicketDefault) DeleteTicket(id int) (err error) {
	return s.rp.DeleteTicket(id)
}

func (s *serviceTicketDefault) GetById(id int) (tkt domain.Ticket, err error) {
	tkt, err = s.rp.GetById(id)
	if err != nil {
		return
	}
	return
}

func (s *serviceTicketDefault) AddTicket(tkt *domain.TicketAttributes) (err error) {
	// add the ticket to the repository
	return s.rp.AddTicket(&domain.TicketAttributes{
		Name:    tkt.Name,
		Email:   tkt.Email,
		Country: tkt.Country,
		Hour:    tkt.Hour,
		Price:   tkt.Price,
	})
}

func (s *serviceTicketDefault) PatchTicket(tkt domain.TicketAttributes, id int) (err error) {
	return s.rp.PatchTicket(tkt, id)
}

func (s *serviceTicketDefault) UpdateTicket(tkt domain.TicketAttributes, id int) (err error) {
	return s.rp.UpdateTicket(tkt, id)
}

func (s *serviceTicketDefault) GetTicketsByDestinationCountry(country string) (t map[int]domain.Ticket, err error) {
	t, err = s.rp.GetTicketsByDestinationCountry(country)
	if err != nil {
		return nil, err
	}
	return
}

func (s *serviceTicketDefault) GetTicketProportion(country string) (propotion float64, err error) {
	return s.rp.GetTicketProportion(country)
}

package loader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/jsdzgzMELI/Desafio-Cierre/internal/domain"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	// loader := &LoaderTicketCSV{}
	// loader.filePath = filePath
	// loader.Load()
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
// REFACTOR TO TICKET
// REFACTOR TO TICKET
// REFACTOR TO TICKET
// REFACTOR TO TICKET
// REFACTOR TO TICKET

func (t *LoaderTicketCSV) Load() (tkt map[int]domain.Ticket, lastId int, err error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	tkt = make(map[int]domain.Ticket)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			return nil, 0, err
		}

		// serialize the record
		id, err := strconv.Atoi(record[0])
		if err != nil {
			err = fmt.Errorf("error parsing id: %v", err)
			return nil, 0, err
		}
		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			err = fmt.Errorf("error parsing price: %v", err)
			return nil, 0, err
		}
		ticket := domain.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}
		tkt[id] = domain.Ticket{
			Id:         id,
			Attributes: ticket,
		}

		// add the ticket to the map
		// tkt[id] = ticket
		lastId = id
	}

	return
}

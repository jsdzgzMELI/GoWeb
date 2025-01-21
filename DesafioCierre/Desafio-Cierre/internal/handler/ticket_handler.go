package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jsdzgzMELI/Desafio-Cierre/internal/domain"
	"github.com/jsdzgzMELI/Desafio-Cierre/internal/service"
	"github.com/jsdzgzMELI/Desafio-Cierre/pkg/web/response"
)

type HandlerTicketDefault struct {
	service service.ServiceTicket
}

func NewHandlerTicketDefault(service service.ServiceTicket) *HandlerTicketDefault {
	return &HandlerTicketDefault{
		service: service,
	}
}

func (th *HandlerTicketDefault) GetHttp(w http.ResponseWriter, r *http.Request) {
	t, err := th.service.GetTickets()
	if err != nil {
		code := http.StatusNotFound
		body := &response.Response{Message: "No tickets found", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &response.Response{Message: "Tickets", Data: t}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (th *HandlerTicketDefault) GetByIdHttp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || id == "0" {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "id is required and can't be 0"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid id"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	t, err := th.service.GetById(intId)
	if err != nil {
		code := http.StatusNotFound
		body := &response.Response{Message: "No tickets found", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &response.Response{Message: "Tickets", Data: t}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (th *HandlerTicketDefault) GetCountryHttp(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")
	if country == "" {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "id is required and can't be 0"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	t, err := th.service.GetTicketsByDestinationCountry(country)
	if err != nil {
		code := http.StatusNotFound
		body := &response.Response{Message: "No tickets found", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &response.Response{Message: "Total ticket amount to the selected destiny are:", Data: len(t)}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (th *HandlerTicketDefault) GetProportionHttp(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")
	if country == "" {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "id is required and can't be 0"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	t, err := th.service.GetTicketProportion(country)
	if err != nil {
		code := http.StatusNotFound
		body := &response.Response{Message: "No tickets found", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &response.Response{Message: "The percentaje of tickets to the selected destiny is:", Data: t}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (th *HandlerTicketDefault) AddHttp(w http.ResponseWriter, r *http.Request) {
	var request domain.TicketAttributes
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &response.Response{Message: "Invalid request body", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	err := th.service.AddTicket(&request)
	if err != nil {
		code := http.StatusInternalServerError
		body := &response.Response{Message: "Error adding ticket", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &response.Response{Message: "Ticket added", Data: request}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (th *HandlerTicketDefault) DeleteHttp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || id == "0" {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "id is required and can't be 0"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid id"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	err = th.service.DeleteTicket(intId)
	if err != nil {
		code := http.StatusNotFound
		body := &response.ErrorResponse{Status: code, Message: "Ticket not found"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	body := &response.Response{Message: "Ticket Deleted", Data: nil}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (th *HandlerTicketDefault) PatchHttp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || id == "0" {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "id is required and can't be 0"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid id"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	var request domain.TicketAttributes

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "error decoding request"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	err = th.service.PatchTicket(request, intId)
	if err != nil {
		code := http.StatusNotFound
		body := &response.Response{Message: err.Error(), Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	t, _ := th.service.GetById(intId)
	code := http.StatusOK
	body := &response.Response{Message: "Ticket Patched", Data: t}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (th *HandlerTicketDefault) UpdateHttp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" || id == "0" {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "id is required and can't be 0"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "Invalid id"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	var request domain.TicketAttributes

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		code := http.StatusBadRequest
		body := &response.ErrorResponse{Status: http.StatusBadRequest, Message: "error decoding request"}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	err = th.service.UpdateTicket(request, intId)
	if err != nil {
		code := http.StatusNotFound
		body := &response.Response{Message: err.Error(), Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	t, _ := th.service.GetById(intId)
	code := http.StatusOK
	body := &response.Response{Message: "Ticket Updated", Data: t}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

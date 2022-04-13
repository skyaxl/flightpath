package api

import (
	"encoding/json"
	"net/http"
)

type FlightService interface {
	CalculateFlighPath(input [][]string) (response []string, err error)
}

type Handler struct {
	service FlightService
}

//NewHandler create new handler
func NewHandler(service FlightService) *Handler {
	return &Handler{service}
}

type Error struct {
	Message string `json:"message,omitempty"`
}

//ServeHTTP to intercept
func (p *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var (
		input [][]string
		err   error
		res   []string
	)

	if err = json.NewDecoder(req.Body).Decode(&input); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(Error{
			Message: err.Error(),
		})
		return
	}

	if res, err = p.service.CalculateFlighPath(input); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(Error{
			Message: err.Error(),
		})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}

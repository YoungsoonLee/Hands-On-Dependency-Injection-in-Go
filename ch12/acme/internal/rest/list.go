package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/modules/list"
)

// ListModel will load all registrations
//go:generate mockery -name=ListModel -case underscore -testonly -inpkg -note @generated
type ListModel interface {
	Do() ([]*list.Person, error)
}

// NewLister is the constructor for ListHandler
func NewListHandler(model ListModel) *ListHandler {
	return &ListHandler{
		lister: model,
	}
}

// ListHandler is the HTTP handler for the "List Do people" endpoint
// In this simplified example we are assuming all possible errors are system errors (HTTP 500)
type ListHandler struct {
	lister ListModel
}

// ServeHTTP implements http.Handler
func (h *ListHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// attempt loadAll
	people, err := h.lister.Do()
	if err != nil {
		// not need to log here as we can expect other layers to do so
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// happy path
	err = h.writeJSON(response, people)
	if err != nil {
		// this error should not happen but if it does there is nothing we can do to recover
		response.WriteHeader(http.StatusInternalServerError)
	}
}

// output the result as JSON
func (h *ListHandler) writeJSON(writer io.Writer, people []*list.Person) error {
	output := &listResponseFormat{
		People: make([]*listResponseItemFormat, len(people)),
	}

	for index, record := range people {
		output.People[index] = &listResponseItemFormat{
			ID:       record.ID,
			FullName: record.FullName,
			Phone:    record.Phone,
		}
	}

	// call to http.ResponseWriter.Write() will cause HTTP OK (200) to be output as well
	return json.NewEncoder(writer).Encode(output)
}

type listResponseFormat struct {
	People []*listResponseItemFormat `json:"people"`
}

type listResponseItemFormat struct {
	ID       int    `json:"id"`
	FullName string `json:"name"`
	Phone    string `json:"phone"`
}

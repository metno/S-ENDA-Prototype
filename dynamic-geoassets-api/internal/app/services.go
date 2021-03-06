package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type ServiceListing []*MetadataService

type MetadataService struct {
	ServiceName string `json:"service_name"`
	ServiceURL  string `json:"service_url"`
	ServiceDoc  string `json:"service_doc"`
	ID          string `json:"id"`
}

func (s *service) serviceCollectionHandler(w http.ResponseWriter, r *http.Request) {
	dataset, err := s.metadataStore.getAllServices()
	if err != nil {
		log.Printf("failed to get dataset: %s", err)
		errorResponse("Failed to access data.", w, r, http.StatusServiceUnavailable)
		return
	}

	payload, err := json.Marshal(dataset)
	if err != nil {
		errorResponse("Failed to serialize response.", w, r, http.StatusInternalServerError)
		return
	}
	okResponse(payload, w, r)
}

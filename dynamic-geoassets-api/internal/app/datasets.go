package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MetadataListing []*MetadataMMD

type MetadataMMD struct {
	ProductName string     `json:"product_name"`
	BoundingBox [4]float32 `json:"bounding_box"`
	Keywords    []string   `json:"keywords"`
	ID          string     `json:"id"`
}

func (s *service) datasetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dataset, err := s.metadataStore.getDataset(params["id"])
	if err != nil {
		log.Printf("failed to get dataset: %s", err)
		serverErrorResponse(err, w, r)
		return
	}

	payload, err := json.Marshal(dataset)
	if err != nil {
		http.Error(w, "Failed to serialize data.", http.StatusInternalServerError)
		return
	}
	okResponse(payload, w, r)
}

func (s *service) putDatasetHandler(w http.ResponseWriter, r *http.Request) {
	payload := r.Body
	defer r.Body.Close()

	dataset := &MetadataMMD{}

	decoder := json.NewDecoder(payload)
	if err := decoder.Decode(dataset); err != nil {
		http.Error(w, "Malformed json content", http.StatusBadRequest)
		return
	}

	if !validDataset(dataset) {
		http.Error(w, "Nonvalid metadata", http.StatusBadRequest)
	}

	storedDataset, err := s.metadataStore.putDataset(dataset)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
	}

	response, err := json.Marshal(storedDataset)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	okResponse(response, w, r)
}

func (s *service) datasetCollectionHandler(w http.ResponseWriter, r *http.Request) {
	listing, err := s.metadataStore.getAllDatasets()
	if err != nil {
		serverErrorResponse(err, w, r)
		return
	}

	payload, err := json.Marshal(listing)
	if err != nil {
		http.Error(w, "Failed to serialize data.", http.StatusInternalServerError)
		return
	}
	okResponse(payload, w, r)
}

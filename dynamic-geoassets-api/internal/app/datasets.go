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
		if err == ErrorDoesNotExist {
			errorResponse("The requested resource does not exist.", w, r, http.StatusNotFound)
			return
		}
		log.Printf("failed to get dataset: %s", err)
		errorResponse("Failed to get dataset.", w, r, http.StatusServiceUnavailable)
		return
	}

	payload, err := json.Marshal(dataset)
	if err != nil {
		errorResponse("Failed to serialize data.", w, r, http.StatusInternalServerError)
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
		errorResponse("Malformed json content", w, r, http.StatusBadRequest)
		return
	}

	if !validDataset(dataset) {
		errorResponse("Nonvalid metadata", w,r, http.StatusBadRequest)
	}

	storedDataset, err := s.metadataStore.putDataset(dataset)
	if err != nil {
		errorResponse("Failed to store dataset. Try again soon.", w, r, http.StatusServiceUnavailable)
	}

	response, err := json.Marshal(storedDataset)
	if err != nil {
		errorResponse("Failed to serialize dataset", w, r, http.StatusInternalServerError)
	}
	okResponse(response, w, r)
}

func (s *service) datasetCollectionHandler(w http.ResponseWriter, r *http.Request) {
	listing, err := s.metadataStore.getAllDatasets()
	if err != nil {
		errorResponse("Failed to access datasets.", w, r, http.StatusServiceUnavailable)
		return
	}

	payload, err := json.Marshal(listing)
	if err != nil {
		errorResponse("Failed to serialize response.", w, r, http.StatusInternalServerError)
		return
	}
	okResponse(payload, w, r)
}

func validDataset(dataset *MetadataMMD) bool {
	return true
}

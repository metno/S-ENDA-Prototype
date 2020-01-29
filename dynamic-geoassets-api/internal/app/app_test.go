package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/metno/S-ENDA-documentation/dynamic-geoassets-api/pkg/metaservice"
)

func testService() *service {
	templates := template.Must(template.ParseGlob("../../templates/*"))

	return NewService(templates, StaticStore)
}

func TestHTTPBasic(t *testing.T) {
	service := testService()

	basicHTTPTests := []struct {
		url        string
		handler    http.HandlerFunc
		statusCode int
	}{
		{"/healthz", metaservice.HealthzHandler(service.checkHealthz), 200},
		{"/", service.docsHandler, 200},
		{"/api/v1/service", service.serviceCollectionHandler, 200},
	}

	for _, httpTest := range basicHTTPTests {
		req := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:8080%s", httpTest.url), nil)
		w := httptest.NewRecorder()
		httpTest.handler(w, req)

		response := w.Result()
		if response.StatusCode != httpTest.statusCode {
			t.Errorf("http handler test for %s:\n got: status code %d\n wanted: %d",
				httpTest.url, response.StatusCode, httpTest.statusCode)
		}
	}
}
func TestHTTPdatasetCollection(t *testing.T) {
	service := testService()

	req := httptest.NewRequest("GET", "http://localhost:8080/api/v1/dataset", nil)
	w := httptest.NewRecorder()
	service.datasetCollectionHandler(w, req)
	resp := w.Result().Body

	decoder := json.NewDecoder(resp)
	var datasets MetadataListing
	err := decoder.Decode(&datasets)
	if err != nil {
		t.Error("got: unable to decode json response, wanted: succefully decoded http json response.")
	}

	if len(datasets) != 2 {
		t.Errorf("\nExpected number of datasets in response:\n got: %d\n wanted: 2", len(datasets))
	}
}

var testDataset = `
{
  "bounding_box": [
    120,
    79,
    -10,
    90
  ],
  "keywords": [
    "Wind",
    "Pressure"
  ],
  "product_name": "Norway forecast 100m supergood"
}
`

func TestHTTPStoreDataset(t *testing.T) {
	service := testService()

	reader := strings.NewReader(testDataset)
	req := httptest.NewRequest("POST", "http://localhost:8080/api/v1/dataset", reader)
	w := httptest.NewRecorder()
	service.putDatasetHandler(w, req)
	resp := w.Result().Body

	decoder := json.NewDecoder(resp)
	var dataset MetadataMMD
	err := decoder.Decode(&dataset)
	if err != nil {
		t.Errorf(" got: unable to decode json response: %s\n wanted: succefully decoded http json response.", err)
	}
}

func TestHTTPGetDataset(t *testing.T) {
	service := testService()

	req := httptest.NewRequest("GET", "http://localhost:8080/api/v1/dataset/587ee038-41ab-11ea-b3e8-3b50360377e9", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "587ee038-41ab-11ea-b3e8-3b50360377e9",
	})
	w := httptest.NewRecorder()

	service.datasetHandler(w, req)

	if w.Result().StatusCode != 200 {
		t.Errorf("Expected status ok:\n got: %d, wanted: 200", w.Result().StatusCode)
		return
	}

	resp := w.Result().Body

	decoder := json.NewDecoder(resp)
	var dataset MetadataMMD
	err := decoder.Decode(&dataset)
	if err != nil {
		t.Errorf(" got: unable to decode json response: %s\n wanted: succefully decoded http json response.", err)
		return
	}

	if dataset.ProductName != "Topaz" {
		t.Errorf("Expected dataset in response:\n got: %s \n wanted: Topaz", dataset.ProductName)
	}
}

// Package app starts up the service and datastore for Geodynamic assets.
package app

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	gorilla "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/metno/S-ENDA-documentation/dynamic-geoassets-api/pkg/metaservice"
	"github.com/metno/S-ENDA-documentation/dynamic-geoassets-api/pkg/middleware"
)

const staticFilesDir = "./static/"

type service struct {
	about          *metaservice.About
	htmlTemplates  *template.Template
	staticFilesDir string
	InternalRouter *mux.Router
	ExternalRouter *mux.Router
	metadataStore  MetadataStore
}

type ServerError struct {
	ErrMsg string `json:"error"`
}

type MetadataStore interface {
	putDataset(md *MetadataMMD) (*MetadataMMD, error)
	getDataset(name string) (*MetadataMMD, error)
	getAllDatasets() (MetadataListing, error)
	getAllServices() (ServiceListing, error)
}

// NewService returns a struct containing that have the routes and handlers for this application.
func NewService(templates *template.Template, metadataStore MetadataStore) *service {
	service := service{
		about:          about(),
		htmlTemplates:  templates,
		staticFilesDir: staticFilesDir,
		InternalRouter: mux.NewRouter(),
		ExternalRouter: mux.NewRouter(),
		metadataStore:  metadataStore,
	}
	service.routes()

	return &service
}

func (s *service) routes() {
	var metrics = middleware.NewServiceMetrics(middleware.MetricsOpts{
		Name:            "sendaregistration",
		Description:     "Metadata registration, editing and listing.",
		ResponseBuckets: []float64{0.001, 0.002, 0.1, 0.5},
	})

	s.ExternalRouter.HandleFunc("/api/v1/dataset/{id}", metrics.Endpoint("/v1/dataset", s.datasetHandler)).Methods("GET")
	s.ExternalRouter.HandleFunc("/api/v1/dataset", metrics.Endpoint("/v1/dataset", s.datasetCollectionHandler)).Methods("GET")
	s.ExternalRouter.HandleFunc("/api/v1/dataset", metrics.Endpoint("/v1/dataset", s.putDatasetHandler)).Methods("POST")

	s.ExternalRouter.HandleFunc("/api/v1/service", metrics.Endpoint("/v1/service", s.serviceCollectionHandler)).Methods("GET")

	// Health of the service
	s.ExternalRouter.HandleFunc("/api/v1/healthz", metaservice.HealthzHandler(s.checkHealthz))

	// Service discovery metadata for the world
	s.ExternalRouter.Handle("/api/v1/about", proxyHeaders(metaservice.AboutHandler(s.about)))

	// Metrics of the service(s) for this app.
	s.InternalRouter.Handle("/metrics", metrics.Handler())

	// Documentation of the service(s)
	s.ExternalRouter.HandleFunc("/docs/{page}", s.docsHandler)

	// Swagger UI
	swui := http.StripPrefix("/swaggerui", http.FileServer(http.Dir("./static/swaggerui/")))
	s.ExternalRouter.PathPrefix("/swaggerui").Handler(swui)

	// Static assets.
	s.ExternalRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(s.staticFilesDir))))

	// Send root path of the http service to the docs index page.
	s.ExternalRouter.HandleFunc("/", s.docsHandler)
}

// proxyHeaders is a http handler middleware function for setting scheme and host correctly when behind a proxy.
// Usually needed when the response consists of urls to the service.
func proxyHeaders(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	setSchemeIfEmpty := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Scheme == "" {
			r.URL.Scheme = "http"
		}
		next(w, r)
	}
	return gorilla.ProxyHeaders(http.HandlerFunc(setSchemeIfEmpty))
}

// html docs generated from templates.
func (s *service) docsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, exists := params["page"]

	var err error
	if exists != true {
		err = s.htmlTemplates.ExecuteTemplate(w, "index", s.about)
	} else {
		err = s.htmlTemplates.ExecuteTemplate(w, page, s.about)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// checkHealthz is supplied to metaservice.HealthzHandler as a callback function.
func (s *service) checkHealthz() (*metaservice.Healthz, error) {
	return &metaservice.Healthz{
		Status:      metaservice.HealthzStatusHealthy,
		Description: "No deps, so everything is ok all the time.",
	}, nil
}

func okResponse(payload []byte, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=60")
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(payload)
	if err != nil {
		log.Printf("could send response to req %q: %s", r.URL, err)
	}
}

func serverErrorResponse(errMsg error, w http.ResponseWriter, r *http.Request) {
	errResponse := ServerError{
		ErrMsg: errMsg.Error(),
	}

	payload, err := json.Marshal(errResponse)
	if err != nil {
		http.Error(w, "Failed to serialize data.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(payload)
	if err != nil {
		log.Printf("could send response to req %q: %s", r.URL, err)
	}
}

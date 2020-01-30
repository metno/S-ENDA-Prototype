package app

import (
	"sync"

	"github.com/google/uuid"
)

// StaticMetadataStore implements the MetadataStore interfac, using static datasets.
type StaticMetadataStore struct {
	mux      sync.Mutex
	datasets map[string]*MetadataMMD
	services map[string]*MetadataService
}

var StaticStore = StaticMetadataStore{
	datasets: map[string]*MetadataMMD{
		"50f599ec-41ab-11ea-b7d5-07e9c6bbe0fb": &MetadataMMD{
			ProductName: "arome arctic",
			ID:          "50f599ec-41ab-11ea-b7d5-07e9c6bbe0fb",
			Keywords:    []string{"Pressure", "Temperature"},
			BoundingBox: [4]float32{20.0, 100.1, 70.1, 119.1},
		},
		"587ee038-41ab-11ea-b3e8-3b50360377e9": &MetadataMMD{
			ProductName: "Topaz",
			ID:          "587ee038-41ab-11ea-b3e8-3b50360377e9",
			Keywords:    []string{"Wave height"},
			BoundingBox: [4]float32{20.0, 100.1, 70.1, 119.1},
		},
	},
	services: map[string]*MetadataService{
		"27abed66-4287-11ea-811b-d3869b6043ba": &MetadataService{
			ServiceName: "OpenDAP best estimate",
			ID:          "27abed66-4287-11ea-811b-d3869b6043ba",
			ServiceURL:  "thredds.met.no/dap-best-estimate/api",
			ServiceDoc:  "thredds.met.no/dap-best-estimate",
		},
	},
}

func (ss StaticMetadataStore) putDataset(dataset *MetadataMMD) (*MetadataMMD, error) {
	uuid := uuid.New()
	dataset.ID = uuid.String()

	ss.mux.Lock()
	defer ss.mux.Unlock()
	ss.datasets[dataset.ID] = dataset

	return dataset, nil
}

func (ss StaticMetadataStore) getAllDatasets() (MetadataListing, error) {
	var listing MetadataListing

	ss.mux.Lock()
	defer ss.mux.Unlock()
	for key := range ss.datasets {
		listing = append(listing, ss.datasets[key])
	}
	return listing, nil
}

func (ss StaticMetadataStore) getDataset(id string) (*MetadataMMD, error) {
	ss.mux.Lock()
	defer ss.mux.Unlock()
	dataset, ok := ss.datasets[id]
	if !ok {
		return nil, ErrorDoesNotExist
	}

	return dataset, nil
}

func (ss StaticMetadataStore) getAllServices() (ServiceListing, error) {
	var listing ServiceListing

	ss.mux.Lock()
	defer ss.mux.Unlock()
	for key := range ss.services {
		listing = append(listing, ss.services[key])
	}
	return listing, nil
}

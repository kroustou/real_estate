package main

import (
	"testing"
)

type MockBackend struct {}

func (*MockBackend) getQuery(url string) []Ad {
	return []Ad{
		Ad{
			Link: "https://test",
			City: "city",
			Region: "region",
			Price: 20,
			Bathrooms: 2,
			Bedrooms: 2,
			M2: 12,
		},
	}
}

type MockStorageBackend struct {}

func (*MockStorageBackend) push(ads []Ad) {
}

func TestGetAds( t *testing.T ) {
	// try without any queries
	ag := AdGetter{
		storageBackend: &MockStorageBackend{},
		backend: &MockBackend{},
	}
	ag.getAds([]string{})
	if len(ag.ads) != 0 {
		t.Error("Got more than 0 results")
	}
	// try with one query
	ag.getAds([]string{"test-query"})
	if len(ag.ads) != 1 {
		t.Error("Got more than 1 results")
	}
}

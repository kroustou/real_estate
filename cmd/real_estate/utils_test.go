package main

import "testing"

func TestGetAds( t *testing.T ) {
	// try without any queries
	var queries []string
	ag := AdGetter{prometheusUrl: "test"}
	ag.getAds(queries)
}

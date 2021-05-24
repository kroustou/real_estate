package main

import "testing"

func TestGetAds( t *testing.T ) {
	// try without any queries
	var queries []string
	getAds(queries)
}

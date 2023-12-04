// main_test.go
package main

import (
	"ascenda-interview/offer"
	"fmt"
	"testing"
)

func BenchmarkOfferFilter(b *testing.B) {
	offerFilter := offer.NewOfferFilter(
		offer.WithInputFileName("generated_offers"),
		offer.WithOutputFileName("output"),
		offer.WithMaxDate(5),
		offer.WithNotEligibleCategories("Hotel"),
	)
	offerFilter.IOderFilter = &offer.OfferFilterJSON{
		OfferFilter: offerFilter,
	}

	err := offerFilter.LoadFile()
	if err != nil {
		fmt.Println("Error load file offers:", err)
		return
	}
	checkinDate := "2019-12-12"
	for i := 0; i < b.N; i++ {
		err := offerFilter.Filter(checkinDate)
		if err != nil {
			fmt.Println("Error filtering offers:", err)
		}
	}
}

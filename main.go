package main

import (
	"ascenda-interview/offer"
	"fmt"
)

func main() {
	offerFilter := offer.NewOfferFilter(
		offer.WithInputFileName("input"),
		offer.WithOutputFileName("output"),
		offer.WithMaxDate(5),
		offer.WithEligibleCategories("Activity", "Restaurant", "Retail"),
		offer.WithNumberOffer(3),
	)
	offerFilter.IOderFilter = &offer.OfferFilterJSON{
		OfferFilter: offerFilter,
	}

	err := offerFilter.LoadFile()
	if err != nil {
		fmt.Println("Error load file offers:", err)
		return
	}

	for {
		// Check-in date
		var checkinDate string

		// Get user input for checkinDate
		fmt.Print("Enter check-in date (YYYY-MM-DD), or press 'x' to exit: ")
		fmt.Scanln(&checkinDate)

		// Check if the user wants to exit
		if checkinDate == "x" {
			fmt.Println("Exiting program.")
			return
		}

		// Filter and get selected offers
		err := offerFilter.Filter(checkinDate)
		if err != nil {
			fmt.Println("Error filtering offers:", err)
		}

	}

}

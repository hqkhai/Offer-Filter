package offer

import (
	"ascenda-interview/dto"
	"encoding/json"
	"os"
	"time"
)

//Concrete Implementation

type OfferFilterJSON struct {
	*OfferFilter
}

func (of *OfferFilterJSON) readFile() error {
	file, err := os.Open(of.InputFileName + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(of)
	if err != nil {
		return err
	}

	return nil
}

func (of *OfferFilterJSON) writeFile(data interface{}) error {
	file, err := os.Create(of.OutputFileName + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func (of *OfferFilterJSON) validateDate(checkinDate, validDate time.Time) bool {
	if checkinDate.Add(time.Duration(of.MaxDate)*24*time.Hour).Compare(validDate) <= 0 {
		return true
	}
	return false
}

func (of *OfferFilterJSON) filter(checkinDate string) (*dto.OffersData, error) {

	filteredOffers := &dto.OffersData{}

	checkinDateTime, err := time.Parse("2006-01-02", checkinDate)
	if err != nil {
		return nil, err
	}

	//Validate and choose valid offer
	for _, offer := range of.OffersData {
		validToDate, err := time.Parse("2006-01-02", offer.ValidTo)
		if err != nil {
			return nil, err
		}

		//case invalid validToDate
		if v := of.validateDate(checkinDateTime, validToDate); !v {
			continue
		}

		//case category is not eligible
		if _, ok := of.EligibleCategories[offer.Category]; !ok {
			// Swap the root of the min heap (arr[0]) with the last element of the unsorted array
			continue
		}

		//case merchants have multiple elements
		if len(offer.Merchants) > 1 {
			//maxMerchantIndex := 0
			min, secondMin := 0, -1
			for i, merchant := range offer.Merchants {
				if merchant.Distance < offer.Merchants[min].Distance {
					min = i
				}
			}
			for i, merchant := range offer.Merchants {
				if i != min {
					secondMin = i
				}
				if secondMin != -1 && merchant.Distance < offer.Merchants[secondMin].Distance {
					secondMin = i
				}
			}
			offer.Merchants = []dto.Merchant{
				offer.Merchants[min],
			}
			if secondMin != -1 {
				offer.Merchants = append(offer.Merchants, offer.Merchants[secondMin])
			}
		}

		filteredOffers.Offers = append(filteredOffers.Offers, offer)

	}

	chosenOffers := heapSort2ClosestMerchant(filteredOffers.Offers, of.NumberOffer)

	return chosenOffers, nil
}

func heapSort2ClosestMerchant(offers []dto.Offer, numberOffer int) *dto.OffersData {
	chosenOffers := &dto.OffersData{}
	n := len(offers)

	//check empty offers
	if n == 0 {
		return chosenOffers
	}

	// Build a min heap by rearranging the array
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(offers, n, i)
	}
	categorySet := make(map[int]bool)

	// Extract elements from the min heap one by one and place them at the end of the array
	for i := n - 1; i >= 0; i-- {
		if _, ok := categorySet[offers[0].Category]; !ok {
			// Swap the root of the min heap (arr[0]) with the last element of the unsorted array
			chosenOffers.Offers = append(chosenOffers.Offers, offers[0])
			categorySet[offers[0].Category] = true
		}
		if len(chosenOffers.Offers) == numberOffer {
			break
		}
		offers[0], offers[i] = offers[i], offers[0]

		// Heapify the reduced min heap to maintain the heap property
		siftDown(offers, i, 0)
	}
	return chosenOffers
}

// siftDown function to restore the min heap property of a subtree rooted with node i
func siftDown(arr []dto.Offer, n, i int) {
	smallest := i    // Initialize smallest as root
	child := 2*i + 1 // Index of left child

	for child < n {
		// If right child exists and is smaller than left child
		if child+1 < n && arr[child+1].Merchants[0].Distance < arr[child].Merchants[0].Distance {
			child++
		}

		// If child is smaller than root
		if arr[child].Merchants[0].Distance < arr[smallest].Merchants[0].Distance {
			// Swap the root with the smaller child
			arr[child], arr[smallest] = arr[smallest], arr[child]

			// Move to the next child
			smallest = child
			child = 2*smallest + 1
		} else {
			break
		}
	}
}

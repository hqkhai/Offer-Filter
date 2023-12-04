package offer

import (
	"ascenda-interview/constant"
	"ascenda-interview/dto"
)

type IOrderFilter interface {
	readFile() error
	writeFile(data interface{}) error
	filter(checkinDate string) (*dto.OffersData, error)
}

// template method
type OfferFilter struct {
	IOderFilter           IOrderFilter
	MaxDate               int
	OffersData            []dto.Offer `json:"offers"`
	InputFileName         string
	OutputFileName        string
	NotEligibleCategories map[int]bool
}

// Functional Option Pattern
type OfferFilterOption func(*OfferFilter)

func WithMaxDate(maxDate int) func(*OfferFilter) {
	return func(of *OfferFilter) {
		of.MaxDate = maxDate
	}
}

func WithNotEligibleCategories(values ...string) func(*OfferFilter) {
	return func(of *OfferFilter) {
		for _, value := range values {
			of.NotEligibleCategories[constant.CategoryMappingId[value]] = true
		}
	}
}

func WithInputFileName(fileName string) func(*OfferFilter) {
	return func(of *OfferFilter) {
		of.InputFileName = fileName
	}
}

func WithOutputFileName(fileName string) func(*OfferFilter) {
	return func(of *OfferFilter) {
		of.OutputFileName = fileName
	}
}

func NewOfferFilter(opts ...OfferFilterOption) *OfferFilter {
	c := &OfferFilter{
		MaxDate:               5,
		InputFileName:         "input",
		OutputFileName:        "output",
		NotEligibleCategories: map[int]bool{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (of *OfferFilter) LoadFile() error {
	if err := of.IOderFilter.readFile(); err != nil {
		return err
	}
	return nil
}

func (of *OfferFilter) Filter(checkinDate string) error {

	result, err := of.IOderFilter.filter(checkinDate)
	if err != nil {
		return err
	}

	if err := of.IOderFilter.writeFile(result); err != nil {
		return err
	}

	return nil
}

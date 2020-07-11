package main

import (
	"cubic-calculator/pkg/calculating"
	"fmt"
	"log"
)

const (
	FirstURL = "/api/products/1"
	Category = "Air Conditioners"
	BaseURL  = "http://wp8m3he1wt.s3-website-ap-southeast-2.amazonaws.com"
)

func main() {

	calculator := calculating.Calculator{
		APIClient: calculating.NewClient(),
		BaseURL:   BaseURL,
	}

	avgCubicWeight, err := calculator.CalculateAvgCubicWeight(FirstURL, Category)
	if err != nil {
		log.Fatalf("cannot get avg cubic weight %+v", err)
	}

	fmt.Printf("The Average Cubic Weight for %s is %f", Category, avgCubicWeight)
}

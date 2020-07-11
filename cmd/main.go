package main

import (
	"cubic-calculator/pkg/calculating"
	"fmt"
	"log"
	"net/url"
	"os"
)

const (
	//SourceURL is the default source URL
	SourceURL = "http://wp8m3he1wt.s3-website-ap-southeast-2.amazonaws.com/api/products/1"
	//Category is the default category
	Category = "Air Conditioners"
)

func main() {
	urlArg := os.Getenv("SOURCE_URL")
	if urlArg == "" {
		urlArg = SourceURL
	}

	categoryArg := os.Getenv("CATEGORY")
	if categoryArg == "" {
		categoryArg = Category
	}

	u, err := url.Parse(urlArg)
	if err != nil {
		log.Fatalf("Invalid URL %+v", err)
	}

	calculator := calculating.Calculator{
		APIClient: calculating.NewClient(),
		BaseURL:   u,
	}

	avgCubicWeight, err := calculator.CalculateAvgCubicWeight(u.Path, categoryArg)
	if err != nil {
		log.Fatalf("cannot get avg cubic weight %+v", err)
	}

	fmt.Printf("The Average Cubic Weight for %s is %f kg", categoryArg, avgCubicWeight)
}

package calculating

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetAvgCubicWeight(t *testing.T) {
	products := []product{
		{
			Category: "Test",
			Size: productSize{
				Height: 1,
				Width:  1,
				Length: 1,
			},
		},
		{
			Category: "Test",
			Size: productSize{
				Height: 2,
				Width:  2,
				Length: 2,
			},
		},
	}
	avg := getAvgCubicWeight(products)
	want := ((float32(1) / 1000000 * 250) + (float32(2*2*2) / 1000000 * 250)) / 2
	if avg != want {
		t.Errorf("Expecting this avg %f, got %f", want, avg)
	}
}

func TestExtractValidProduct(t *testing.T) {
	products := []product{
		{
			Category: "RightCategory",
			Size: productSize{
				Height: -1,
				Width:  1,
				Length: 1,
			},
		},
		{
			Category: "RightCategory",
			Size: productSize{
				Height: 2,
				Width:  2,
				Length: 2,
			},
		},
		{
			Category: "WrongCategory",
			Size: productSize{
				Height: 3,
				Width:  3,
				Length: 4,
			},
		},
	}
	prds := extractValidProduct(products, "RightCategory")

	if len(prds) != 1 {
		t.Errorf("Expecting to get %d product, got %d", 1, len(prds))
	}
}

func TestGetProducts(t *testing.T) {
	client := APIClient{
		doHTTPReq: func(req *http.Request) (resp *http.Response, err error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(validResponse1))),
			}, nil
		},
	}

	calculator := Calculator{
		APIClient: &client,
	}

	nextURL, prds, err := calculator.getProducts("testURL", "Air Conditioners")
	if err != nil {
		t.Errorf("Error while getting the products %+v", err)
	}
	if len(prds) != 2 {
		t.Errorf("Expecting to get %d product, got %d", 2, len(prds))
	}
	want := "/api/products/2"
	if nextURL != want {
		t.Errorf("Expecting to get URL %s, got %s", want, nextURL)
	}
}

func TestCalculateAvgCubicWeight(t *testing.T) {
	client := APIClient{
		doHTTPReq: func(req *http.Request) (resp *http.Response, err error) {
			if req.URL.Path == "/testURL" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(validResponse1))),
				}, nil
			}
			if req.URL.Path == "/api/products/2" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(validResponse2))),
				}, nil
			}
			return &http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("Will be an error if it gets to here"))),
			}, nil

		},
	}

	calculator := Calculator{
		APIClient: &client,
		BaseURL:   "http://localhost",
	}

	want := ((float32(26*26*5) / 1000000 * 250) + (float32(49.6*38.7*89) / 1000000 * 250) + (float32(49.6*38.7*50) / 1000000 * 250)) / 3
	avg, err := calculator.CalculateAvgCubicWeight("/testURL", "Air Conditioners")
	if err != nil {
		t.Errorf("Error while getting the products %+v", err)
	}
	if avg != want {
		t.Errorf("Expecting this avg %f, got %f", want, avg)
	}
}

var validResponse1 = `
    {
		"objects": [
			{
				"category": "Gadgets", 
				"title": "10 Pack Family Car Sticker Decals", 
				"weight": 120.0, 
				"size": {
					"width": 15.0, 
					"length": 13.0, 
					"height": 1.0
				}
			}, 
			{
				"category": "Air Conditioners", 
				"title": "Window Seal for Portable Air Conditioner Outlets", 
				"weight": 235.0, 
				"size": {
					"width": 26.0, 
					"length": 26.0, 
					"height": 5.0
				}
			}, 
			{
				"category": "Batteries", 
				"title": "10 Pack Kogan CR2032 3V Button Cell Battery", 
				"weight": 60.0, 
				"size": {
					"width": 5.8, 
					"length": 19.0, 
					"height": 0.3
				}
			}, 
			{
				"category": "Cables & Adapters", 
				"title": "3 Pack Apple MFI Certified Lightning to USB Cable (3m)", 
				"weight": 90.0, 
				"size": {
					"width": 10.0, 
					"length": 20.0, 
					"height": 3.0
				}
			}, 
			{
				"category": "Air Conditioners", 
				"title": "Kogan 10,000 BTU Portable Air Conditioner (2.9KW)", 
				"weight": 26200.0, 
				"size": {
					"width": 49.6, 
					"length": 38.7, 
					"height": 89.0
				}
			}
		], 
		"next": "/api/products/2"
	
}`

var validResponse2 = `
    {
		"objects": [				
			{
				"category": "Air Conditioners", 
				"title": "Kogan 10,000 BTU Portable Air Conditioner (2.9KW)", 
				"weight": 26200.0, 
				"size": {
					"width": 49.6, 
					"length": 38.7, 
					"height": 50.0
				}
			}
		], 
		"next": ""
}`

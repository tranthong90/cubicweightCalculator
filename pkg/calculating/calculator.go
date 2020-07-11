package calculating

import (
	"encoding/json"
	"net/http"
	"net/url"
)

//Calculator calculate average cubic weight
type Calculator struct {
	APIClient *APIClient
	BaseURL   *url.URL
}

type apiResponse struct {
	Objects []product `json:"objects"`
	NextURL string    `json:"next"`
}
type product struct {
	Category string      `json:"category"`
	Size     productSize `json:"size"`
}

type productSize struct {
	Width  float32 `json:"width"`
	Length float32 `json:"length"`
	Height float32 `json:"height"`
}

func (ps productSize) IsValid() bool {
	if ps.Width > 0 && ps.Length > 0 && ps.Height > 0 {
		return true
	}
	return false
}

func (ps productSize) GetCubicSize() float32 {
	return (ps.Height * ps.Width * ps.Length) / 1000000 * 250
}

//CalculateAvgCubicWeight calculate avg cubic weight
func (c Calculator) CalculateAvgCubicWeight(urlPath, category string) (float32, error) {
	var products []product
	next := urlPath
	var err error
	var prds []product
	for {
		next, prds, err = c.getProducts(next, category)
		if err != nil {
			return 0, err
		}
		products = append(products, prds...)
		if next == "" {
			break
		}
	}
	return getAvgCubicWeight(products), nil
}

func (c Calculator) getProducts(urlPath, category string) (nextURL string, products []product, err error) {
	c.BaseURL.Path = urlPath

	req, err := http.NewRequest("GET", c.BaseURL.String(), nil)
	if err != nil {
		return "", []product{}, err
	}
	resp, err := c.APIClient.doHTTPReq(req)
	if err != nil {
		return "", []product{}, err
	}
	defer resp.Body.Close()
	var data apiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", []product{}, err
	}

	return data.NextURL, extractValidProduct(data.Objects, category), nil

}

func extractValidProduct(products []product, category string) []product {
	var result []product
	for _, p := range products {
		if p.Size.IsValid() && p.Category == category {
			result = append(result, p)
		}
	}
	return result
}

func getAvgCubicWeight(products []product) float32 {
	sumCubicWeight := float32(0)
	for _, p := range products {
		sumCubicWeight += p.Size.GetCubicSize()
	}
	return sumCubicWeight / float32(len(products))
}

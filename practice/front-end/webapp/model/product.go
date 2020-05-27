package model

import (
	"fmt"
	"sort"
	"strings"
)

// Product :
type Product struct {
	Name             string
	DescriptionShort string
	DescriptionLong  string
	PricePerLiter    float32
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	ImageURL         string
	ID               int
	CategoryID       int
}

// GetProductsForCategory : returns the list of products for the specific category
func GetProductsForCategory(categoryID int) []Product {
	result := []Product{}
	for _, p := range products {
		if p.CategoryID == categoryID {
			result = append(result, p)
		}
	}
	return result
}

// GetProduct : returns product for given productID
func GetProduct(productID int) (*Product, error) {
	for _, p := range products {
		if p.ID == productID {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("Product not found with ID %v", productID)
}

// SearchForProduct : gets products containing the given query string
func SearchForProduct(q string) []Product {
	fp := make(map[Product]int)
	sortedProducts := make([]Product, 0)
	for i := 0; i < len(products); i++ {
		p := products[i]
		numt := strings.Count(strings.ToLower(p.Name), q)
		numt += strings.Count(strings.ToLower(p.DescriptionLong), q)
		fp[p] = numt
		if numt > 0 {
			sortedProducts = append(sortedProducts, p)
		}
	}
	sort.Slice(sortedProducts, func(i, j int) bool {
		return fp[sortedProducts[i]] > fp[sortedProducts[j]]
	})
	return sortedProducts
}

var products []Product = []Product{
	{
		Name:             "Lemon Juice",
		DescriptionShort: "Made from fresh, organic California lemons.",
		DescriptionLong: `Made from premium, organic Meyer lemons. These fruit are left on the tree until they reach the peak of ripeness and then juiced within 8 hours of being picked.
			<br/>
			Lemonade made from our premium juice is sure to make your stand the most popular in the neighborhood.`,
		PricePerLiter:   1.09,
		PricePer10Liter: 1.04,
		Origin:          "California",
		IsOrganic:       true,
		ImageURL:        "lemon.png",
		ID:              1,
		CategoryID:      2,
	}, {
		Name:             "Apple Juice",
		DescriptionShort: "The perfect blend of Washington apples.",
		DescriptionLong:  "The perfect blend of Washington apples.",
		PricePerLiter:    0.89,
		PricePer10Liter:  0.85,
		Origin:           "Ohio",
		IsOrganic:        true,
		ImageURL:         "apple.png",
		ID:               2,
		CategoryID:       1,
	}, {
		Name:             "Watermelon Juice",
		DescriptionShort: "From sun-drenched fields in Florida.",
		DescriptionLong:  "From sun-drenched fields in Florida.",
		PricePerLiter:    3.99,
		PricePer10Liter:  3.84,
		Origin:           "Florida",
		IsOrganic:        true,
		ImageURL:         "watermelon.png",
		ID:               3,
		CategoryID:       1,
	}, {
		Name:             "Kiwi Juice",
		DescriptionShort: "California sunshine and rain distilled into sweet goodness",
		DescriptionLong:  "California sunshine and rain distilled into sweet goodness",
		PricePerLiter:    5.99,
		PricePer10Liter:  5.54,
		Origin:           "California",
		IsOrganic:        false,
		ImageURL:         "kiwi.png",
		ID:               4,
		CategoryID:       1,
	}, {
		Name:             "Mangosteen Juice",
		DescriptionShort: "Our latest taste sensation, imported directly from Hawaii",
		DescriptionLong:  "Our latest taste sensation, imported directly from Hawaii",
		PricePerLiter:    6.87,
		PricePer10Liter:  6.79,
		Origin:           "Hawaii",
		IsOrganic:        false,
		ImageURL:         "mangosteen.png",
		ID:               5,
		CategoryID:       2,
	}, {
		Name:             "Orange Juice",
		DescriptionShort: "Fresh squeezed from Florida's best oranges.",
		DescriptionLong:  "Fresh squeezed from Florida's best oranges.",
		PricePerLiter:    1.67,
		PricePer10Liter:  1.63,
		Origin:           "Florida",
		IsOrganic:        false,
		ImageURL:         "orange.png",
		ID:               6,
		CategoryID:       1,
	}, {
		Name:             "Pineapple Juice",
		DescriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
		DescriptionLong:  "An exotic and refreshing offering. Straight from Hawaii.",
		PricePerLiter:    2.48,
		PricePer10Liter:  2.27,
		Origin:           "Hawaii",
		IsOrganic:        false,
		ImageURL:         "pineapple.png",
		ID:               7,
		CategoryID:       1,
	}, {
		Name:             "Strawberry Juice",
		DescriptionShort: "MThe perfect balance of sweet and tart.",
		DescriptionLong:  "The perfect balance of sweet and tart.",
		PricePerLiter:    4.36,
		PricePer10Liter:  4.27,
		Origin:           "California",
		IsOrganic:        false,
		ImageURL:         "strawberry.png",
		ID:               8,
		CategoryID:       2,
	},
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

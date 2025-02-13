// Code generated by github.com/john-markham/gqlgen, DO NOT EDIT.

package model

type Manufacturer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Manufacturer) IsEntity() {}

type Product struct {
	ID           string        `json:"id"`
	Manufacturer *Manufacturer `json:"manufacturer"`
	Upc          string        `json:"upc"`
	Name         string        `json:"name"`
	Price        int           `json:"price"`
}

func (Product) IsEntity() {}

type Query struct {
}

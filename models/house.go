package models

type House struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HousePoints int    `json:"points"`
	LogoURL     string `json:"logo_url"`
}

package models

type House struct {
	ID          uint   `json:"id"          gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HousePoints int    `json:"points"      gorm:"column:house_points"`
	LogoURL     string `json:"logo_url"    gorm:"column:logo_url"`
}

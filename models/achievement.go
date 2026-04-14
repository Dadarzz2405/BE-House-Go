package models

type Achievement struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	HouseID int    `json:"house_id"`
	House   *House `json:"house"`
}

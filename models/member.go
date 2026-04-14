package models

type Member struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	HouseID int    `json:"house_id"`
	House   *House `json:"house"`
}

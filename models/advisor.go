package models

type Advisor struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	Bio          string `json:"bio"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`

	HouseID int    `json:"house_id"`
	House   *House `json:"house"`
}

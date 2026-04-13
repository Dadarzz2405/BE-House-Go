package models

type Advisor struct {
	ID           int
	Name         string
	Role         string
	Bio          string
	Username     string
	PasswordHash string

	HouseID int
	House   *House
}

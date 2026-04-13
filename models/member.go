package models

type Member struct {
	ID   int
	Name string
	Role string

	HouseID int
	House   *House
}

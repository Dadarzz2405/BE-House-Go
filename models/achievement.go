package models

type Achievement struct {
	ID          int
	Name        string
	Description string

	HouseID int
	House   *House
}

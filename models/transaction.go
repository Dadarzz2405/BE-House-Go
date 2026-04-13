package models

import "time"

type PointTransaction struct {
	ID           int
	PointsChange int
	Reason       string
	Timestamp    time.Time

	HouseID int
	AdminID *int

	House *House
	Admin *Admin
}

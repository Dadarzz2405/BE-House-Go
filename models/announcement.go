package models

import "time"

type Announcement struct {
	ID        int
	Title     string
	Content   string
	ImageURL  *string
	CreatedAt time.Time

	HouseID   int
	CaptainID *int

	House   *House
	Captain *Captain
}

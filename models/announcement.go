package models

import "time"

type Announcement struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageURL  *string   `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	HouseID   int       `json:"house_id"`
	CaptainID *int      `json:"captain_id"`
	House     *House    `json:"house"`
	Captain   *Captain  `json:"captain"`
}

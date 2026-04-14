package models

type Captain struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`

	HouseID int    `json:"house_id"`
	House   *House `json:"house"`

	// Relations
	Announcements []Announcement `json:"announcements"`
}

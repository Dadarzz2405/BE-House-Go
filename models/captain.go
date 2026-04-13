package models

type Captain struct {
	ID           int
	Name         string
	Username     string
	PasswordHash string

	HouseID int
	House   *House

	// Relations
	Announcements []Announcement
}

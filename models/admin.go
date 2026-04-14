package models

type Admin struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	// Relations
	PointTransactions []PointTransaction `json:"point_transactions"`
}

package models

type Admin struct {
	ID           int
	Name         string
	Username     string
	PasswordHash string
	// Relations
	PointTransactions []PointTransaction
}

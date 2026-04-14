package models

import "time"

type PointTransaction struct {
	ID           int       `json:"id"`
	PointsChange int       `json:"points_change"`
	Reason       string    `json:"reason"`
	Timestamp    time.Time `json:"timestamp"`

	HouseID int  `json:"house_id"`
	AdminID *int `json:"admin_id"`

	House *House `json:"house"`
	Admin *Admin `json:"admin"`
}

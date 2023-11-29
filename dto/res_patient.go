package dto

import "time"

type ResPatient struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	Indication string    `json:"indication"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

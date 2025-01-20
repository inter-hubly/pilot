package base

import "time"

type Entity struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RemovedAt time.Time `json:"removed_at"`
	RemovedBy string    `json:"removed_by"`
	Removed   bool      `json:"removed"`
}

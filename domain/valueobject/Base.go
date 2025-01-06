package valueobject

import "time"

type BaseObject struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Delete    bool      `json:"delete"`
}

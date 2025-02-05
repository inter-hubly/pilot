package entity

import (
	"github.com/inter-hubly/pilot/domain/base"
)

type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	LoginAttempt uint8  `json:"loginAttempt"`
	base.Entity  `json:",inline"`
}

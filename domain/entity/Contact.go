package entity

import (
	"github.com/inter-hubly/pilot/domain/base"
	"github.com/inter-hubly/pilot/domain/valueobject"
	"github.com/pkg/errors"
)

type Contact struct {
	Id          string                             `json:"id" bson:"_id,omitempty"`
	Name        string                             `json:"name" bson:"name"`
	Phone       string                             `json:"phone" bson:"phone"`
	Variables   []valueobject.Pair[string, string] `json:"variables" bson:"variables"`
	base.Entity `bson:",inline"`
}

func (c *Contact) GetVariableByName(name string) (string, error) {
	for _, v := range c.Variables {
		if v.Key == name {
			return v.Value, nil
		}
	}
	return "", errors.New("variable not found")
}

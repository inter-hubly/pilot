package entity

import (
	"github.com/inter-hubly/pilot/domain/base"
	"github.com/inter-hubly/pilot/domain/valueobject"
)

type Campaign struct {
	Id          string                             `json:"id" bson:"_id,omitempty"`
	Name        string                             `json:"name" bson:"name"`
	Template    base.TemplateInfo                  `json:"template" bson:"template"`
	ContactsId  []string                           `json:"contactsId" bson:"contactsId"`
	Variables   []valueobject.Pair[string, string] `json:"variables" bson:"variables"`
	base.Entity `bson:",inline"`
}

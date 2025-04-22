package entity

import (
	"github.com/inter-hubly/pilot/domain/base"
	"github.com/inter-hubly/pilot/domain/valueobject"
)

type Campaign struct {
	Id          string                             `json:"id" bson:"_id,omitempty"`
	Name        string                             `json:"name" bson:"name"`
	Flows       map[string]*Flow                   `json:"flows" bson:"flows"`
	Template    base.TemplateInfo                  `json:"template" bson:"template"`
	ContactsId  []string                           `json:"contactsId" bson:"contactsId"`
	IaContext   string                             `json:"iaContext" bson:"iaContext"`
	Variables   []valueobject.Pair[string, string] `json:"variables" bson:"variables"`
	base.Entity `bson:",inline"`
}

type Flow struct {
	Id              string `json:"id" bson:"_id,omitempty"`
	Role            string `json:"role" bson:"role,omitempty"`
	Message         string `json:"message" bson:"message"`
	IsIaInteraction bool   `json:"isIaInteraction" bson:"isIaInteraction"`
}

package base

import (
	"github.com/inter-hubly/pilot/domain/valueobject"
)

type StartTemplateDto struct {
	To               string                             `json:"to"`
	CampaignId       string                             `json:"campaignId"`
	Template         TemplateInfo                       `json:"templateInfo"`
	HasIaInteraction bool                               `json:"hasIaInteraction"`
	Parameters       []valueobject.Pair[string, string] `json:"parameters"`
}

type SendTextDto struct {
	To      string `json:"to"`
	Message string `json:"message"`
	IsOwner bool   `json:"isOwner,omitempty"`
}

type TemplateInfo struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

package entity

import (
	"github.com/inter-hubly/pilot/domain/base"
)

type Client struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	AppId         string `json:"appId"`
	PhoneNumberId string `json:"phoneNumberId"`
	BusinessId    string `json:"businessId"`
	AccessToken   string `json:"-"`
	base.Entity   `json:",inline"`
}

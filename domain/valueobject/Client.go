package valueobject

type Client struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	AppId         string `json:"appId"`
	PhoneNumberId string `json:"phoneNumberId"`
	BusinessId    string `json:"businessId"`
	AccessToken   string `json:"-"`
	BaseObject
}

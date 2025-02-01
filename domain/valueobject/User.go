package valueobject

type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ClientId     string `json:"clientId"`
	LoginAttempt uint8  `json:"loginAttempt"`
	TenantId     string `json:"tenantId"`
	BaseObject
}

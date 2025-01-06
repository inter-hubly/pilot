package valueobject

type User struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ClientId     uint64 `json:"clientId"`
	LoginAttempt uint8  `json:"loginAttempt"`
	TenantId     string `json:"tenantId"`
	BaseObject
}

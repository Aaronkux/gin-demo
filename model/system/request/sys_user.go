package request

type Register struct {
	Email        string   `json:"email"`
	Phone        *string  `json:"phone"`
	Password     string   `json:"password"`
	Nickname     string   `json:"nickname"`
	Avatar       string   `json:"avatar"`
	AuthorityIds []string `json:"authorityIds"`
}

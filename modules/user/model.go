package user

type User struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	FirstName       string `json:"firstName"`
}

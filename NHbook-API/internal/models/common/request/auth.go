package request

type Register struct {
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

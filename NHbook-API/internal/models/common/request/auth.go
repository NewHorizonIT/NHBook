package request

type Register struct {
	UserName    string `json:"username" required:"true"`
	Email       string `json:"email" required:"true"`
	Password    string `json:"password" required:"true"`
	PhoneNumber string `json:"phone_number" required:"true"`
}

type Login struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

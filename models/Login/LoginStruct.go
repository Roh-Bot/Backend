package Login

type LoginStruct struct {
	Email       string `json:"email" validate:"email"`
	Mobile      int    `json:"mobile" `
	Password    string `json:"password" validate:"required"`
	Remember_me bool   `json:"remember_me"`
}

type LoginResponse struct {
	StatusCode int
	Error      map[string]string
	Data       map[string]string
}

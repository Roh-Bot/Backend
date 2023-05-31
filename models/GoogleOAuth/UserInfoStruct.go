package GoogleOAuth

type UserInfoStruct struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Verified_Email bool   `json:"verified_email"`
	Name           string `json:"name"`
	Given_Name     string `json:"given_name"`
	Family_Name    string `json:"family_name"`
	Picture        string `json:"picture"`
	Locale         string `json:"locale"`
}

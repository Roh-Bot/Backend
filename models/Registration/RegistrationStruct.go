package Registration

import "time"

type RegistrationStruct struct {
	First_name     string    `json:"first_Name" validate:"required,min=2,max=14"`
	Middle_name    string    `json:"middle_name" validate:"min=2,max=14"`
	Last_name      string    `json:"last_name" validate:"required,min=2,max=14"`
	Password       string    `json:"password" validate:"required,min=8"`
	Email          string    `json:"email" validate:"email,min=4"`
	Phone_no       int       `json:"phone_no" validate:"required,min=10"`
	Dob            time.Time `json:"dob"`
	Address_line_1 string    `json:"address_line_1" validate:"required,min=1"`
	Address_line_2 string    `json:"address_line_2" validate:"required,min=1"`
	City_id        int       `json:"city_id"`
	State_id       int       `json:"state_id"`
	Pincode        int       `json:"pincode" validate:"required"`
	Referred_by    int       `json:"referred_by"`
	Reference_code int       `json:"reference_code"`
}

type RegistrationResponse struct {
	StatusCode int
	Error      map[string]string
	Data       map[string]string
}

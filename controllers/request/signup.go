package request

type SignUpRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	Password    string `json:"password"`
}

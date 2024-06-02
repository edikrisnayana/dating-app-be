package request

type CreateAccountRequest struct {
	Email       string
	PhoneNumber string
	FirstName   string
	LastName    string
	Gender      string
	Password    string
}

package response

type ResponseStatus string

const (
	SUCCESS ResponseStatus = "SUCCESS"
	USER_NOT_EXIST ResponseStatus = "USER_NOT_EXIST"
	USER_ALREADY_EXIST ResponseStatus = "USER_ALREADY_EXIST"
)
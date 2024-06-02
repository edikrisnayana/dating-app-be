package response

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error *MessageResponse `json:"error"`
}

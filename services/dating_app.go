package services

import (
	"datingAppBE/accessors"
	request "datingAppBE/services/request"
	"datingAppBE/services/response"
)

type AccountService interface {
	init(*accessors.AccountEntityAccessor) AccountService
	CreateAccount(*request.CreateAccountRequest) *response.CreateAccountResponse
}

type accountService struct {
	accountEntityAccessor accessors.AccountEntityAccessor
}

func CreateService(accountEntityAccessor *accessors.AccountEntityAccessor) AccountService {
	var service AccountService = new(accountService)
	return service.init(accountEntityAccessor)
}

func (service *accountService) init(accountEntityAccessor *accessors.AccountEntityAccessor) AccountService {
	service.accountEntityAccessor = *accountEntityAccessor
	return service
}

func (service *accountService) CreateAccount(request *request.CreateAccountRequest) *response.CreateAccountResponse {
	return nil
}

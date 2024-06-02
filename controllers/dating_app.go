package controllers

import (
	crequest "datingAppBE/controllers/request"
	response "datingAppBE/controllers/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DatingAppController interface {
	init() DatingAppController
	SignUp(*gin.Context)
}

type datingAppController struct {
}

func CreateDatingAppController() DatingAppController {
	var controller DatingAppController = new(datingAppController)
	return controller.init()
}

func (controller *datingAppController) init() DatingAppController {
	return controller
}

func (controller *datingAppController) SignUp(ctx *gin.Context) {
	defer handleError(ctx)

	var request crequest.SignUpRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		log.Printf("binding error: %s", err.Error())
		ctx.IndentedJSON(http.StatusBadRequest, response.ErrorResponse{Error: &response.MessageResponse{Message: err.Error()}})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, response.ErrorResponse{Error: nil})
}

package controllers

import (
	"datingAppBE/controllers/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func handleError(ctx *gin.Context) {
	if err := recover(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%v", err)})
	}
}

func AuthCheck(functions ...gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer handleError(ctx)

		authHeader := ctx.Request.Header["Authorization"]
		if len(authHeader) < 1 {
			ctx.IndentedJSON(http.StatusUnauthorized, response.ErrorResponse{&response.MessageResponse{Message: "Request not authorized"}})
			return
		}

		auth := strings.Fields(authHeader[0])
		if len(auth) != 2 || auth[0] != "Bearer" || auth[1] != "authdata" {
			ctx.IndentedJSON(http.StatusUnauthorized, response.ErrorResponse{&response.MessageResponse{Message: "Request not authorized"}})
			return
		}

		for _, f := range functions {
			f(ctx)
		}
	}
}

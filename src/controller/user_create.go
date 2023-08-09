package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmmisss/golang-default/src/configuration/logger"
	"github.com/jimmmisss/golang-default/src/configuration/validation"
	"github.com/jimmmisss/golang-default/src/controller/model/request"
	"github.com/jimmmisss/golang-default/src/controller/model/response"
	"go.uber.org/zap"
	"net/http"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init Createuser controller", zap.String("journey", "createdUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createdUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	response := response.UserResponse{
		Id:    "TEST",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully", zap.String("journey", "createdUser"))

	c.JSON(http.StatusOK, response)
}

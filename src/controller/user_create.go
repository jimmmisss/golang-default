package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	restErr "github.com/jimmmisss/golang-default/src/configuration/rest_err"
	"github.com/jimmmisss/golang-default/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := restErr.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Print(userRequest)
}

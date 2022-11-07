package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/response"
)

type User struct {
	Id       int
	Username string
	Sex      uint8
	Mobile   string
	Avatar   string
	Birthday string
}

func (u User) Get(c *gin.Context) {

	a := User{
		Id:       0,
		Username: "",
		Sex:      0,
		Mobile:   "",
		Avatar:   "",
		Birthday: "",
	}

	response.Success(c, a)
	return
}

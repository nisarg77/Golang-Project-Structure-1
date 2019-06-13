package v1Service

import (
	u "Golang-Project-Structure-1/apiHelpers"
	"Golang-Project-Structure-1/models"
	res "Golang-Project-Structure-1/resources/api/v1"
)

type UserService struct {
	User models.User
}

func (us *UserService) UserList() map[string]interface{} {
	user := us.User

	userData := res.UserResponse{
		ID:    user.ID,
		Name:  "test",
		Email: "test@gmail.com",
	}
	response := u.Message(0, "This is from version 1 api")
	response["data"] = userData
	return response
}

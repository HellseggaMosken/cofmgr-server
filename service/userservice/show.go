package userservice

import (
	"cofmgr/model"
	"cofmgr/service"
	"cofmgr/service/serializer"
)

func Show(id string) (*serializer.User, service.Status) {
	user := model.FindUserWithID(id)
	if user == nil {
		return nil, service.StatusNotFoundItem
	}
	return serializer.BuildUser(user), service.StatusOK
}

func List() ([]*serializer.User, service.Status) {
	users := model.ListUsers()
	return serializer.BuildUsers(users), service.StatusOK
}

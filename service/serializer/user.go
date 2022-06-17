package serializer

import "cofmgr/model"

type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	MidName    string `json:"midName"`
	LastName   string `json:"lastName"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Department string `json:"department"`
}

// BuildUser return nil when the param is nil.
func BuildUser(user *model.User) *User {
	if user == nil {
		return nil
	}
	return &User{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		MidName:    user.MidName,
		LastName:   user.LastName,
		City:       user.City,
		Country:    user.Country,
		Department: user.Department,
	}
}

// BuildUsers return nil when the param is nil.
func BuildUsers(users []model.User) []*User {
	res := make([]*User, len(users))
	for i := range users {
		res[i] = BuildUser(&users[i])
	}
	return res
}

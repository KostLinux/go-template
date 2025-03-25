package response

import "go-template/model"

// UserMapper converts a single user model to a user response
func UserMapper(user *model.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

// UserListMapper converts a slice of user models to a user list response
func UserListMapper(users []model.User) UserList {
	responses := make([]User, len(users))
	for i := range users {
		responses[i] = UserMapper(&users[i])
	}
	return UserList{Users: responses}
}

// SingleUserMapper converts a user pointer to a UserList with one user
func SingleUserMapper(user *model.User) UserList {
	return UserList{
		Users: []User{UserMapper(user)},
	}
}

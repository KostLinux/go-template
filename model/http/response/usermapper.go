package response

import "go-template/model"

// ToResponse converts a user model to a user response
func ToResponse(user *model.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToResponseList(users []model.User) UserList {
	responses := make([]User, 0, len(users))
	for _, user := range users {
		responses = append(responses, ToResponse(&user))
	}
	return UserList{Users: responses}
}

package dto

import "github.com/sherwin-77/go-tix/internal/entity"

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	ID       string `param:"id" validate:"required,uuid"`
	Email    string `json:"email" validate:"omitempty,email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID       string          `json:"id"`
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Roles    []*RoleResponse `json:"roles,omitempty"`
}

func NewUserResponse(user *entity.User) UserResponse {
	userResponse := UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}
	for _, role := range user.Roles {
		roleResponse := NewRoleResponse(role)
		userResponse.Roles = append(userResponse.Roles, &roleResponse)
	}

	return userResponse
}

func NewUsersResponse(users []entity.User) []UserResponse {
	var res []UserResponse
	for _, user := range users {
		res = append(res, NewUserResponse(&user))
	}
	return res
}

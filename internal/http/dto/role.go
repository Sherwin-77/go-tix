package dto

import "github.com/sherwin-77/go-tix/internal/entity"

type RoleRequest struct {
	Name      string `json:"name" validate:"required"`
	AuthLevel int    `json:"auth_level" validate:"required"`
}

type UpdateRoleRequest struct {
	RoleRequest
	ID string `param:"id" validate:"required,uuid"`
}

type ChangeRoleRequest struct {
	UserID string                  `param:"id" validate:"required,uuid"`
	Items  []ChangeRoleRequestItem `json:"items" validate:"required"`
}

type ChangeRoleRequestItem struct {
	ID     string `json:"id" validate:"required,uuid"`
	Action string `json:"action" validate:"required,oneof=add remove"`
}

type RoleResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AuthLevel int    `json:"auth_level"`
}

func NewRoleResponse(role *entity.Role) RoleResponse {
	return RoleResponse{
		ID:        role.ID.String(),
		Name:      role.Name,
		AuthLevel: role.AuthLevel,
	}
}

func NewRolesResponse(roles []entity.Role) []RoleResponse {
	var response []RoleResponse
	for _, role := range roles {
		response = append(response, NewRoleResponse(&role))
	}
	return response
}

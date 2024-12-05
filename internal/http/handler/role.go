package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/http/dto"
	"github.com/sherwin-77/go-tix/internal/service"
	"github.com/sherwin-77/go-tix/pkg/response"
	"net/http"
)

type RoleHandler struct {
	RoleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) RoleHandler {
	return RoleHandler{roleService}
}

// GetRoles
//
//	@Summary	Get All Roles
//	@Tags		[Admin] Role
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.Response{data=[]dto.RoleResponse}
//	@Router		/admin/roles [get]
func (h *RoleHandler) GetRoles(ctx echo.Context) error {
	roles, err := h.RoleService.GetRoles(ctx.Request().Context())

	if err != nil {
		return err
	}

	rolesResponse := dto.NewRolesResponse(roles)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", rolesResponse, nil))
}

// GetRoleByID
//
//	@Summary	Get Role By ID
//	@Tags		[Admin] Role
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"The Role ID"
//	@Success	200	{object}	response.Response{data=dto.RoleResponse}
//	@Router		/admin/roles/{id} [get]
func (h *RoleHandler) GetRoleByID(ctx echo.Context) error {
	roleID := ctx.Param("id")
	if roleID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	role, err := h.RoleService.GetRoleByID(ctx.Request().Context(), roleID)
	if err != nil {
		return err
	}

	roleResponse := dto.NewRoleResponse(role)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", roleResponse, nil))
}

// CreateRole
//
//	@Summary	Create Role
//	@Tags		[Admin] Role
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.RoleRequest	true	"Role Request"
//	@Success	201		{object}	response.Response{data=dto.RoleResponse}
//	@Router		/admin/roles [post]
func (h *RoleHandler) CreateRole(ctx echo.Context) error {
	var req dto.RoleRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	role, err := h.RoleService.CreateRole(ctx.Request().Context(), req)

	if err != nil {
		return err
	}

	roleResponse := dto.NewRoleResponse(role)

	return ctx.JSON(http.StatusCreated, response.NewResponse(http.StatusCreated, "Role Created", roleResponse, nil))
}

// UpdateRole
//
//	@Summary	Update Role
//	@Tags		[Admin] Role
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string					true	"The Role ID"
//	@Param		request	body		dto.UpdateRoleRequest	true	"Update Role Request"
//	@Success	200		{object}	response.Response{data=dto.RoleResponse}
//	@Router		/admin/roles/{id} [patch]
func (h *RoleHandler) UpdateRole(ctx echo.Context) error {
	var req dto.UpdateRoleRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	role, err := h.RoleService.UpdateRole(ctx.Request().Context(), req)

	if err != nil {
		return err
	}

	roleResponse := dto.NewRoleResponse(role)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Role Updated", roleResponse, nil))
}

// DeleteRole
//
//	@Summary	Delete Role
//	@Tags		[Admin] Role
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"The Role ID"
//	@Success	200	{object}	response.Response{data=nil}
//	@Router		/admin/roles/{id} [delete]
func (h *RoleHandler) DeleteRole(ctx echo.Context) error {
	roleID := ctx.Param("id")
	if roleID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	if err := h.RoleService.DeleteRole(ctx.Request().Context(), roleID); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Role Deleted", nil, nil))
}

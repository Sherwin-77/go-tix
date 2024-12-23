package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/internal/http/dto"
	"github.com/sherwin-77/go-tix/internal/service"
	"github.com/sherwin-77/go-tix/pkg/response"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService}
}

/**
 * Admin Handlers
**/

// GetUsers
//
//	@summary	Get All Users
//	@tags		User
//	@accept		json
//	@produce	json
//	@Success	200	{object}	response.Response{data=[]dto.UserResponse}
//	@router		/admin/users [get]
func (h *UserHandler) GetUsers(ctx echo.Context) error {
	users, meta, err := h.userService.GetUsers(ctx.Request().Context(), ctx.QueryParams())
	if err != nil {
		return err
	}

	usersResponse := dto.NewUsersResponse(users)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", usersResponse, meta))
}

// GetUserByID
//
//	@Summary	Get User By ID
//	@Tags		[Admin] User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"The User ID"
//	@Success	200	{object}	response.Response{data=dto.UserResponse}
//	@Router		/admin/users/{id} [get]
func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	userID := ctx.Param("id")
	if userID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	user, err := h.userService.GetUserByID(ctx.Request().Context(), userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", user, nil))
}

// CreateUser
//
//	@Summary	Create User
//	@Tags		[Admin] User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.UserRequest	true	"User Request"
//	@Success	201		{object}	response.Response{data=dto.UserResponse}
//	@Router		/admin/users [post]
func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var req dto.UserRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	user, isFirstUser, err := h.userService.Register(ctx.Request().Context(), req)

	if err != nil {
		return err
	}

	userResponse := dto.NewUserResponse(user)
	msg := "User Created"
	if isFirstUser {
		msg += ". Because this is the first user, admin role has been assigned"
	}
	return ctx.JSON(http.StatusCreated, response.NewResponse(http.StatusCreated, msg, userResponse, nil))
}

// UpdateUser
//
//	@Summary	Update User
//	@Tags		[Admin] User
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string					true	"The User ID"
//	@Param		request	body		dto.UpdateUserRequest	true	"Update User Request"
//	@Success	200		{object}	response.Response{data=dto.UserResponse}
//	@Router		/admin/users/{id} [patch]
func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var req dto.UpdateUserRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	user, err := h.userService.UpdateUser(ctx.Request().Context(), req)

	if err != nil {
		return err
	}

	userResponse := dto.NewUserResponse(user)

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "User Updated", userResponse, nil))
}

// DeleteUser
//
//	@Summary	Delete User
//	@Tags		[Admin] User
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"The User ID"
//	@Success	200	{object}	response.Response{data=dto.UserResponse}
//	@Router		/admin/users/{id} [delete]
func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	userID := ctx.Param("id")
	if userID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	if err := h.userService.DeleteUser(ctx.Request().Context(), userID); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "User Deleted", nil, nil))
}

// ChangeRole
//
//	@Summary	Change Role
//	@Tags		[Admin] User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.ChangeRoleRequest	true	"Change Role Request"
//	@Success	200		{object}	response.Response{data=nil}
//	@Router		/users/role [patch]
func (h *UserHandler) ChangeRole(ctx echo.Context) error {
	var req dto.ChangeRoleRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	if err := h.userService.ChangeRole(ctx.Request().Context(), req); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Role Changed", nil, nil))
}

/**
 * User Handlers
**/

// Register
//
//	@Summary	Register User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.UserRequest	true	"User Request"
//	@Success	201		{object}	response.Response{data=dto.UserResponse}
//	@Router		/register [post]
func (h *UserHandler) Register(ctx echo.Context) error {
	return h.CreateUser(ctx)
}

// Login
//
//	@Summary	Login User
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.LoginRequest				true	"Login Request"
//	@Success	200		{object}	response.Response{data=string}	"token"
//	@Router		/login [post]
func (h *UserHandler) Login(ctx echo.Context) error {
	var req dto.LoginRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	token, err := h.userService.Login(ctx.Request().Context(), req)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Login Success", token, nil))
}

// ShowProfile
//
//	@Summary	Show Profile
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.Response{data=dto.UserResponse}
func (h *UserHandler) ShowProfile(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	ctx.SetParamNames("id")
	ctx.SetParamValues(userID)
	return h.GetUserByID(ctx)
}

// EditProfile
//
//	@Summary	Edit Profile
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.UpdateUserRequest	true	"Update User Request"
//	@Success	200		{object}	response.Response{data=dto.UserResponse}
//	@Router		/profile [patch]
func (h *UserHandler) EditProfile(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	// Inject user_id
	ctx.SetParamNames("id")
	ctx.SetParamValues(userID)

	var req dto.UpdateUserRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	return h.UpdateUser(ctx)
}

package handler_test

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/configs"
	"github.com/sherwin-77/go-tix/internal/entity"
	"github.com/sherwin-77/go-tix/internal/http/handler"
	mock_service "github.com/sherwin-77/go-tix/test/mock/service"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type RoleTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	roleService *mock_service.MockRoleService
	roleHandler handler.RoleHandler
}

func (s *RoleTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.roleService = mock_service.NewMockRoleService(s.ctrl)
	s.roleHandler = handler.NewRoleHandler(s.roleService)
}

func TestRole(t *testing.T) {
	suite.Run(t, new(RoleTestSuite))
}

func (s *RoleTestSuite) TestGetRoles() {
	roles := make([]entity.Role, 0)

	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/roles", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.roleService.EXPECT().GetRoles(c.Request().Context()).Return(nil, errorTest)
		err := s.roleHandler.GetRoles(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/roles", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.roleService.EXPECT().GetRoles(c.Request().Context()).Return(roles, nil)
		err := s.roleHandler.GetRoles(c)

		s.Nil(err)
		s.Equal(http.StatusOK, rec.Code)
	})
}

func (s *RoleTestSuite) TestGetRoleByID() {
	role := &entity.Role{}

	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/roles", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		s.roleService.EXPECT().GetRoleByID(c.Request().Context(), "1").Return(nil, errorTest)
		err := s.roleHandler.GetRoleByID(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/roles", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		s.roleService.EXPECT().GetRoleByID(c.Request().Context(), "1").Return(role, nil)
		err := s.roleHandler.GetRoleByID(c)

		s.Nil(err)
		s.Equal(http.StatusOK, rec.Code)
	})
}

func (s *RoleTestSuite) TestCreateRole() {
	s.Run("Error binding", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/roles", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := s.roleHandler.CreateRole(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error validating", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/roles", strings.NewReader(`{"name": ""}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := s.roleHandler.CreateRole(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/roles", strings.NewReader(`{"name": "test", "auth_level": 1}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.roleService.EXPECT().CreateRole(c.Request().Context(), gomock.Any()).Return(nil, errorTest)
		err := s.roleHandler.CreateRole(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/roles", strings.NewReader(`{"name": "test", "auth_level": 1}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.roleService.EXPECT().CreateRole(c.Request().Context(), gomock.Any()).Return(&entity.Role{}, nil)
		err := s.roleHandler.CreateRole(c)

		s.Nil(err)
		s.Equal(http.StatusCreated, rec.Code)
	})
}

func (s *RoleTestSuite) TestUpdateRole() {
	s.Run("Error binding", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/roles", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		err := s.roleHandler.UpdateRole(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error validating", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/roles", strings.NewReader(`{"name": "test"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		err := s.roleHandler.UpdateRole(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/roles", strings.NewReader(`{"name": "test", "auth_level": 1}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		s.roleService.EXPECT().UpdateRole(c.Request().Context(), gomock.Any()).Return(nil, errorTest)
		err := s.roleHandler.UpdateRole(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/roles", strings.NewReader(`{"name": "test", "auth_level": 1}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		s.roleService.EXPECT().UpdateRole(c.Request().Context(), gomock.Any()).Return(&entity.Role{}, nil)
		err := s.roleHandler.UpdateRole(c)

		s.Nil(err)
		s.Equal(http.StatusOK, rec.Code)
	})
}

func (s *RoleTestSuite) TestDeleteRole() {
	roleID := uuid.NewString()

	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/roles", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(roleID)

		s.roleService.EXPECT().DeleteRole(c.Request().Context(), roleID).Return(errorTest)
		err := s.roleHandler.DeleteRole(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/roles", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(roleID)

		s.roleService.EXPECT().DeleteRole(c.Request().Context(), roleID).Return(nil)
		err := s.roleHandler.DeleteRole(c)

		s.Nil(err)
		s.Equal(http.StatusOK, rec.Code)
	})
}

package handler_test

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-echo-template/configs"
	"github.com/sherwin-77/go-echo-template/internal/entity"
	"github.com/sherwin-77/go-echo-template/internal/http/handler"
	mock_service "github.com/sherwin-77/go-echo-template/test/mock/service"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	userService *mock_service.MockUserService
	userHandler handler.UserHandler
}

func (s *UserTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.userService = mock_service.NewMockUserService(s.ctrl)
	s.userHandler = handler.NewUserHandler(s.userService)
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestGetUsers() {
	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.userService.EXPECT().GetUsers(c.Request().Context(), gomock.Any()).Return(nil, nil, errorTest)
		err := s.userHandler.GetUsers(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.userService.EXPECT().GetUsers(c.Request().Context(), gomock.Any()).Return(nil, nil, nil)
		err := s.userHandler.GetUsers(c)

		s.Nil(err)
		s.Equal(http.StatusOK, rec.Code)
	})
}

func (s *UserTestSuite) TestGetUserByID() {
	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		s.userService.EXPECT().GetUserByID(c.Request().Context(), "1").Return(nil, errorTest)
		err := s.userHandler.GetUserByID(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		s.userService.EXPECT().GetUserByID(c.Request().Context(), "1").Return(nil, nil)
		err := s.userHandler.GetUserByID(c)

		s.Nil(err)
		s.Equal(http.StatusOK, rec.Code)
	})
}

func (s *UserTestSuite) TestRegisterUser() {
	s.Run("Error binding", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := s.userHandler.Register(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error validating", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name": "test"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := s.userHandler.Register(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"username": "test", "email": "test@example.com", "password": "test"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.userService.EXPECT().Register(c.Request().Context(), gomock.Any()).Return(nil, false, errorTest)
		err := s.userHandler.Register(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"username": "test", "email": "test@example.com", "password": "test"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.userService.EXPECT().Register(c.Request().Context(), gomock.Any()).Return(&entity.User{}, true, nil)
		err := s.userHandler.Register(c)

		s.Nil(err)
		s.Equal(http.StatusCreated, rec.Code)
	})
}

func (s *UserTestSuite) TestLoginUser() {
	s.Run("Error binding", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/users/login", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := s.userHandler.Login(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error validating", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(`{"email": "invalid"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := s.userHandler.Login(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error service", func() {
		errorTest := errors.New("error")

		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(`{"email": "test@example.com", "password": "test"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.userService.EXPECT().Login(c.Request().Context(), gomock.Any()).Return("", errorTest)
		err := s.userHandler.Login(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(`{"email": "test@example.com", "password": "test"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.userService.EXPECT().Login(c.Request().Context(), gomock.Any()).Return("token", nil)
		err := s.userHandler.Login(c)

		s.Nil(err)

		body := rec.Body.String()
		var res map[string]interface{}
		err = json.Unmarshal([]byte(body), &res)
		s.Nil(err)

		s.Equal("token", res["data"])
		s.Equal(http.StatusOK, rec.Code)
	})
}

func (s *UserTestSuite) TestEditProfile() {
	s.Run("Error binding", func() {
		userID := uuid.NewString()
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user_id", userID)

		err := s.userHandler.EditProfile(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error validate", func() {
		userID := uuid.NewString()
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/profile", strings.NewReader(`{"email": "invalid"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user_id", userID)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(userID)

		err := s.userHandler.EditProfile(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})
}

func (s *UserTestSuite) TestUpdateUser() {
	s.Run("Error binding", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := s.userHandler.UpdateUser(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error validating", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/users", strings.NewReader(`{"email": "invalid"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		err := s.userHandler.UpdateUser(c)

		s.ErrorAs(err, &validator.ValidationErrors{})
	})

	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/users", strings.NewReader(`{"email": "test@example.com"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		s.userService.EXPECT().UpdateUser(c.Request().Context(), gomock.Any()).Return(&entity.User{}, errorTest)
		err := s.userHandler.UpdateUser(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		e.Validator = configs.NewAppValidator()
		req := httptest.NewRequest(http.MethodPatch, "/users", strings.NewReader(`{"email": "test@example.com"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		s.userService.EXPECT().UpdateUser(c.Request().Context(), gomock.Any()).Return(&entity.User{}, nil)
		err := s.userHandler.UpdateUser(c)

		s.Nil(err)
		s.Equal(http.StatusOK, rec.Code)
	})
}

func (s *UserTestSuite) TestDeleteUser() {
	s.Run("Error service", func() {
		errorTest := errors.New("error")
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		s.userService.EXPECT().DeleteUser(c.Request().Context(), gomock.Any()).Return(errorTest)
		err := s.userHandler.DeleteUser(c)

		s.ErrorIs(err, errorTest)
	})

	s.Run("Success", func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(uuid.NewString())

		s.userService.EXPECT().DeleteUser(c.Request().Context(), gomock.Any()).Return(nil)
		err := s.userHandler.DeleteUser(c)

		s.Nil(err)
		s.Equal(http.StatusOK, rec.Code)
	})
}

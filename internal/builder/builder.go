package builder

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-tix/configs"
	"github.com/sherwin-77/go-tix/internal/http/handler"
	"github.com/sherwin-77/go-tix/internal/http/middlewares"
	"github.com/sherwin-77/go-tix/internal/http/router"
	"github.com/sherwin-77/go-tix/internal/repository"
	"github.com/sherwin-77/go-tix/internal/service"
	"github.com/sherwin-77/go-tix/pkg/caches"
	"github.com/sherwin-77/go-tix/pkg/constants"
	"github.com/sherwin-77/go-tix/pkg/query"
	"github.com/sherwin-77/go-tix/pkg/tokens"
	"gorm.io/gorm"
)

func BuildV1Routes(config *configs.Config, db *gorm.DB, cache caches.Cache, group *echo.Group) {
	g := group.Group("/v1")

	// Initialize middlewares
	middleware := middlewares.NewMiddleware()
	authMiddleware := middlewares.NewAuthMiddleware(config, db)

	// Initialize repositories
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)

	// Initialize builders
	userBuilder := query.NewBuilder(
		[]query.FilterParam{
			{DisplayName: "Email", FieldName: "email", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
			{DisplayName: "Name", FieldName: "name", InternalName: "username", DisplayFilterType: constants.FilterResponsePartialText, FilterType: query.FilterTypePartial},
		},
		[]query.SortParam{
			{DisplayName: "Email", FieldName: "email", InternalName: "Email"},
			{DisplayName: "Username", FieldName: "username", InternalName: "username"},
		},
		query.SortParam{DisplayName: "Email", FieldName: "email", InternalName: "email", Direction: query.SortDirectionAscending},
	)

	// Initialize services
	tokenService := tokens.NewTokenService(config.JWTSecret)
	userService := service.NewUserService(tokenService, userRepository, roleRepository, userBuilder, cache)
	roleService := service.NewRoleService(roleRepository, cache)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)

	// Register routes
	userRoutes, userMiddlewares := router.UserRoutes(userHandler, authMiddleware)
	for _, route := range userRoutes {
		m := append(userMiddlewares, route.Middlewares...)
		g.Add(route.Method, route.Path, route.Handler, m...)
	}

	adminGroup := g.Group("/admin")

	adminRoutes, adminMiddlewares := router.AdminRoutes(userHandler, roleHandler, middleware, authMiddleware)
	for _, route := range adminRoutes {
		m := append(adminMiddlewares, route.Middlewares...)
		adminGroup.Add(route.Method, route.Path, route.Handler, m...)
	}
}

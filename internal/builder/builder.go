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
	eventRepository := repository.NewEventRepository(db)

	// Initialize builders
	userBuilder := NewUserQueryBuilder()
	eventBuilder := NewEventQueryBuilder()

	// Initialize services
	tokenService := tokens.NewTokenService(config.JWTSecret)
	userService := service.NewUserService(tokenService, userRepository, roleRepository, userBuilder, cache)
	roleService := service.NewRoleService(roleRepository, cache)
	eventService := service.NewEventService(eventRepository, eventBuilder, cache)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)
	eventHandler := handler.NewEventHandler(eventService)

	// Register routes
	userRoutes, userMiddlewares := router.UserRoutes(
		userHandler,
		eventHandler,
		authMiddleware,
		middleware,
	)
	for _, route := range userRoutes {
		m := append(userMiddlewares, route.Middlewares...)
		g.Add(route.Method, route.Path, route.Handler, m...)
	}

	adminGroup := g.Group("/admin")

	adminRoutes, adminMiddlewares := router.AdminRoutes(
		userHandler,
		roleHandler,
		eventHandler,
		middleware,
		authMiddleware,
	)
	for _, route := range adminRoutes {
		m := append(adminMiddlewares, route.Middlewares...)
		adminGroup.Add(route.Method, route.Path, route.Handler, m...)
	}
}

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
	snapPaymentRepository := repository.NewSnapPaymentRepository(db)
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	eventRepository := repository.NewEventRepository(db)
	eventApprovalRepository := repository.NewEventApprovalRepository(db)
	ticketRepository := repository.NewTicketRepository(db)
	saleInvoiceRepository := repository.NewSaleInvoiceRepository(db)

	// Initialize builders
	userBuilder := NewUserQueryBuilder()
	eventBuilder := NewEventQueryBuilder()
	eventApprovalBuilder := NewEventApprovalQueryBuilder()
	saleInvoiceBuilder := NewSaleInvoiceQueryBuilder()

	// Initialize services
	tokenService := tokens.NewTokenService(config.JWTSecret)
	midtransService := service.NewMidtransService(config.Midtrans)
	transactionService := service.NewTransactionService(midtransService, snapPaymentRepository)

	userService := service.NewUserService(tokenService, userRepository, roleRepository, userBuilder, cache)
	roleService := service.NewRoleService(roleRepository, cache)
	eventService := service.NewEventService(userRepository, eventRepository, eventApprovalRepository, eventBuilder, cache)
	eventApprovalServices := service.NewEventApprovalService(eventService, eventApprovalRepository, eventApprovalBuilder)
	saleInvoiceServices := service.NewSaleInvoiceService(saleInvoiceRepository, ticketRepository, transactionService, saleInvoiceBuilder)

	webhookService := service.NewWebhookService(config.Midtrans, saleInvoiceRepository, snapPaymentRepository)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)
	eventHandler := handler.NewEventHandler(eventService)
	eventApprovalHandler := handler.NewEventApprovalHandler(eventApprovalServices)
	saleInvoiceHandler := handler.NewSaleInvoiceHandler(saleInvoiceServices)

	webhookHandler := handler.NewWebhookHandler(webhookService)

	// Register routes
	userRoutes, userMiddlewares := router.UserRoutes(
		userHandler,
		eventHandler,
		saleInvoiceHandler,
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
		eventApprovalHandler,
		middleware,
		authMiddleware,
	)
	for _, route := range adminRoutes {
		m := append(adminMiddlewares, route.Middlewares...)
		adminGroup.Add(route.Method, route.Path, route.Handler, m...)
	}

	webhookGroup := g.Group("/webhook")

	webhookRoutes, webhookMiddlewares := router.WebhookRoutes(webhookHandler)
	for _, route := range webhookRoutes {
		m := append(webhookMiddlewares, route.Middlewares...)
		webhookGroup.Add(route.Method, route.Path, route.Handler, m...)
	}
}

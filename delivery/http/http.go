package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
	"user-service/delivery/http/healthcheck"
	"user-service/delivery/http/user"
	"user-service/delivery/http/wallet"
	middleware2 "user-service/middleware"
	"user-service/usecase"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	var (
		e = echo.New()
	)

	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 10 * time.Second,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		//AllowOriginFunc: func(origin string) (bool, error) {
		//	return regexp.MatchString(
		//		`^https:\/\/(|[a-zA-Z0-9]+[a-zA-Z0-9-._]*[a-zA-Z0-9]+\.)teqnological.asia$`,
		//		origin,
		//	)
		//},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))

	// Health check use for microservice
	healthcheck.Init(e.Group("/health-check"))
	// API docs
	//if !config.GetConfig().Stage.IsProd() {
	//	e.GET("/docs/*", echoSwagger.WrapHandler)
	//}

	// APIs
	api := e.Group("/api")
	user.Init(api.Group("/user"), useCase)
	wallet.Init(api.Group("/wallets", middleware2.VerifyToken), useCase)

	return e
}

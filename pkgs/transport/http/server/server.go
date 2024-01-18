package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trekkstay/pkgs/log"
	"trekkstay/pkgs/transport/http/middleware"
	"trekkstay/pkgs/transport/http/route"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "trekkstay/docs"
)

type Option func(*HTTPServer)

type HTTPServer struct {
	Name                    string
	Port                    int
	StrictSlash             bool
	Routes                  []route.Route
	GroupRoutes             []route.GroupRoute
	Middlewares             []func(c *gin.Context)
	GinOptions              []route.GinOption
	GracefulShutdownTimeout time.Duration
	OnCloseFunc             func()
}

func NewHTTPServer(options ...Option) *HTTPServer {
	s := &HTTPServer{}
	for _, option := range options {
		option(s)
	}
	return s
}

func (s *HTTPServer) Run() {
	gin.SetMode(gin.DebugMode)

	// Request ID middleware
	s.Middlewares = append(s.Middlewares, middleware.RequestID())
	s.Middlewares = append(s.Middlewares, middleware.Recover())

	// Setup route
	r := route.NewGin(
		route.AddMiddlewares(s.Middlewares...),
		route.AddHealthCheckRoute(),
		route.AddNoRouteHandler(),
		route.StrictSlash(s.StrictSlash),
		route.AddGroupRoutes(s.GroupRoutes),
		route.AddRoutes(s.Routes),
		route.AddGinOptions(s.GinOptions...),
	)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	hs := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: r,
	}

	// Graceful shutdown
	idleConnectionClosed := make(chan struct{})
	go func() {
		cs := make(chan os.Signal, 1)
		signal.Notify(cs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		<-cs

		ctx, cancel := context.WithTimeout(context.Background(), s.GracefulShutdownTimeout)
		defer cancel()

		log.JsonLogger.Info("Server is shutting down")

		if err := hs.Shutdown(ctx); err != nil {
			log.JsonLogger.Error(err.Error())
		}

		s.OnCloseFunc()

		<-ctx.Done()
		close(idleConnectionClosed)
	}()

	go func() {
		color.Green("➯ http server started on :%d", s.Port)
		if err := hs.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.JsonLogger.Error(err.Error())
			os.Exit(1)
		} else {
			log.JsonLogger.Info("Server has graceful shutdown completely")
		}
	}()

	<-idleConnectionClosed
}

func AddName(n string) Option {
	return func(s *HTTPServer) {
		s.Name = n
	}
}

func AddPort(p int) Option {
	return func(s *HTTPServer) {
		s.Port = p
	}
}

func AddMiddlewares(m []func(c *gin.Context)) Option {
	return func(s *HTTPServer) {
		s.Middlewares = append(s.Middlewares, m...)
	}
}

func AddOnCloseFunc(f func()) Option {
	return func(s *HTTPServer) {
		s.OnCloseFunc = f
	}
}

func AddGinOptions(o []route.GinOption) Option {
	return func(s *HTTPServer) {
		s.GinOptions = o
	}
}

func (s *HTTPServer) AddRoutes(r []route.Route) {
	s.Routes = r
}

func (s *HTTPServer) AddGroupRoutes(gr []route.GroupRoute) {
	s.GroupRoutes = gr
}

func StrictSlash() Option {
	return func(s *HTTPServer) {
		s.StrictSlash = true
	}
}

func SetGracefulShutdownTimeout(d time.Duration) Option {
	return func(s *HTTPServer) {
		s.GracefulShutdownTimeout = d
	}
}

func MustRun(s *HTTPServer, err error) {
	if err != nil {
		panic(err)
	}
	s.Run()
}

package api

import (
	endpoints "invokes/internal/api/endpoints"
	handlers "invokes/internal/api/handlers"
	"invokes/internal/utils"
	"strings"

	docs "invokes/cmd/invokes/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"
)

type Router struct {
	RouterGin *gin.Engine

	Env *handlers.Env
}

func (r *Router) Initialize(env *handlers.Env) {

	r.Env = env

	r.GinInitialize()
}

func (r *Router) Run() error {

	var err error

	err = r.GinRun()

	return err
}

func (r *Router) GinInitialize() {
	if utils.Logger.GetLevel() <= logrus.WarnLevel {
		gin.SetMode(gin.ReleaseMode)
	}

	r.RouterGin = gin.New()

	r.RouterGin.Use(ginlogrus.Logger(utils.Logger), gin.Recovery())

	r.GinRoutesBinding()
}

func (r *Router) GinRun() error {

	r.RouterGin.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(r.Env.Config.Cors.AllowedOrigins, ","),
		AllowMethods: strings.Split(r.Env.Config.Cors.AllowedMethods, ","),
		AllowHeaders: strings.Split(r.Env.Config.Cors.AllowedHeaders, ","),
	}))

	return r.RouterGin.Run(r.Env.Config.Port)

	return nil
}

func (r *Router) GinRoutesBinding() {

	// Swagger
	docs.SwaggerInfo.BasePath = r.Env.Config.Basepath
	r.RouterGin.Handle("GET", "/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Default routes
	r.RouterGin.Handle("GET", "ping", endpoints.Ping(r.Env))
	r.RouterGin.Handle("GET", "version", endpoints.GetVersion(r.Env))

	// Users routes
	r.RouterGin.Handle("GET", "users", endpoints.GetUsers(r.Env))

	// Invoice routes
	r.RouterGin.Handle("GET", "invoices", endpoints.GetInvoices(r.Env))
	r.RouterGin.Handle("POST", "invoice", endpoints.PostInvoice(r.Env))

	// Transaction routes
	r.RouterGin.Handle("POST", "transaction", endpoints.PostTransaction(r.Env))
}

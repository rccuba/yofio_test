package app

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"test_robert_yofio/docs"
	"test_robert_yofio/internal/endpoint"
	creditAssignment "test_robert_yofio/internal/function"
	"test_robert_yofio/internal/middleware"
	"test_robert_yofio/internal/repository"
	"test_robert_yofio/internal/service"
	"test_robert_yofio/internal/static"
)

type App struct {
	Router               *mux.Router
	Logg                 log.Logger
	InvestmentEndpoints  endpoint.InvestmentEndpoints
	InvestmentService    service.InvestmentService
	InvestmentRepository repository.InvestmentRepository
	InvestmentMiddleware middleware.InvestmentMiddleware
}

func (a *App) Run(addr string) error {
	err := http.ListenAndServe(addr, a.Router)
	return err
}

func (a *App) Initialize() (err error) {
	fmt.Println(static.MsgResponseStartApplication)
	muxObj := mux.NewRouter()
	muxObj.Use(middleware.CORS)
	a.Router = muxObj
	values := []interface{}{static.KeyType, static.SUCCESS, static.KeyURL, static.URLStartingNow, static.KeyMessage, static.MsgResponseStartingNow}
	middleware.LoggingOperation(a.Logg, values...)
	a.InitializeEndpoints()
	a.InitializeRoutes()
	a.InitializeSwagger()
	return err
}

// routing
func (a *App) InitializeRoutes() {
	var options []httptransport.ServerOption
	a.Router.PathPrefix(static.URLApi).Handler(httpSwagger.WrapHandler)
	a.Router.Methods(http.MethodPost).Path(static.URLCreditAssignment).Handler(a.CreditAssignment(options))
	values := []interface{}{static.KeyType, static.SUCCESS, static.KeyURL, static.URLStartingNow, static.KeyMessage, static.MsgResponseStartingRoutes}
	middleware.LoggingOperation(a.Logg, values...)
}

//swagger
func (a *App) InitializeSwagger() {
	docs.SwaggerInfo.Title = static.MsgApiRestTitle
	docs.SwaggerInfo.Description = static.MsgApiRestDescription
	docs.SwaggerInfo.Version = static.MsgApiRestVersion1
	docs.SwaggerInfo.Host = viper.GetString(static.APP_HOST) + ":" + viper.GetString(static.APP_PORT)
	docs.SwaggerInfo.BasePath = static.URLStartingNow
	docs.SwaggerInfo.Schemes = []string{static.SchemaHttp}
	values := []interface{}{static.KeyType, static.SUCCESS, static.KeyURL, static.URLStartingNow, static.KeyMessage, static.MsgResponseStartingSwagger}
	middleware.LoggingOperation(a.Logg, values...)
}

// ENDPOINTS
func (a *App) InitializeEndpoints() {
	creditAssigner := creditAssignment.NewCreditAssigner()
	a.InvestmentRepository = repository.NewInvestmentRepository(creditAssigner)
	a.InvestmentService = service.NewInvestmentService(a.InvestmentRepository, a.Logg)
	a.InvestmentMiddleware = middleware.NewInvestmentMiddleware(a.Logg)
	a.InvestmentEndpoints = endpoint.MakeInvestmentEndpoints(&a.InvestmentService, a.InvestmentMiddleware)
	values := []interface{}{static.KeyType, static.SUCCESS, static.KeyURL, static.URLStartingNow, static.KeyMessage, static.MsgResponseStartingEndpoints}
	middleware.LoggingOperation(a.Logg, values...)
}

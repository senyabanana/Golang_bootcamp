// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"flag"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"strconv"

	"candy-server/restapi/operations"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

//go:generate swagger generate server --target ../../ex00 --name CandyServer --spec ../candy-api.yaml --principal interface{}

func configureFlags(api *operations.CandyServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CandyServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.BuyCandyHandler == nil {
		api.BuyCandyHandler = operations.BuyCandyHandlerFunc(func(params operations.BuyCandyParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.BuyCandy has not yet been implemented")
		})
	}

	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(func(params operations.BuyCandyParams) middleware.Responder {
		// Получаем данные из запроса
		money := *params.Order.Money
		candyType := *params.Order.CandyType
		candyCount := *params.Order.CandyCount

		// Цены на конфеты
		candyPrices := map[string]int{
			"CE": 10,
			"AA": 15,
			"NT": 17,
			"DE": 21,
			"YR": 23,
		}

		// Проверка на корректность типа конфет
		price, validCandy := candyPrices[candyType]
		if !validCandy || candyCount <= 0 {
			return operations.NewBuyCandyBadRequest().WithPayload(&operations.BuyCandyBadRequestBody{
				Error: "Invalid candy type or count",
			})
		}

		// Считаем общую стоимость конфет
		totalCost := price * int(candyCount)

		// Проверка, достаточно ли денег
		if money < int64(totalCost) {
			return operations.NewBuyCandyPaymentRequired().WithPayload(&operations.BuyCandyPaymentRequiredBody{
				Error: "You need " + strconv.Itoa(totalCost-int(money)) + " more money!",
			})
		}

		// Если денег достаточно, считаем сдачу
		change := int(money) - totalCost

		// Возвращаем успешный ответ
		return operations.NewBuyCandyCreated().WithPayload(&operations.BuyCandyCreatedBody{
			Thanks: "Thank you!",
			Change: int64(change),
		})
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	portFlag := flag.String("port", "3333", "Port to run the server on")
	flag.Parse()

	s.Addr = ":" + *portFlag
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

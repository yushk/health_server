// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	mws "github.com/yushk/health_backend/apiserver/middlewares"
	"github.com/yushk/health_backend/apiserver/restapi/operations"
	"github.com/yushk/health_backend/apiserver/restapi/operations/oauth"
	"github.com/yushk/health_backend/apiserver/restapi/operations/user"
	"github.com/yushk/health_backend/apiserver/server"
	v1 "github.com/yushk/health_backend/apiserver/v1"
)

//go:generate swagger generate server --target ../../apiserver --name Health --spec ../swagger/swagger.yaml --model-package v1 --principal v1.Principal

func configureFlags(api *operations.HealthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HealthAPI) http.Handler {
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
	api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.OAuth2Auth = func(token string, scopes []string) (*v1.Principal, error) {
		return server.RoleAuth(token, scopes)
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// user.LoginMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// user.ModifyUserPasswordMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// user.ResetPasswordMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// oauth.TokenMaxParseMemory = 32 << 20

	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *v1.Principal) middleware.Responder {
		return server.CreateUser(params, principal)
	})
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteUser(params, principal)
	})
	api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams, principal *v1.Principal) middleware.Responder {
		return server.GetUser(params, principal)
	})
	api.UserGetUserInfoHandler = user.GetUserInfoHandlerFunc(func(params user.GetUserInfoParams, principal *v1.Principal) middleware.Responder {
		return server.GetUserInfo(params, principal)
	})
	api.UserGetUsersHandler = user.GetUsersHandlerFunc(func(params user.GetUsersParams, principal *v1.Principal) middleware.Responder {
		return server.GetUsers(params, principal)
	})
	api.UserLoginHandler = user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
		return server.Login(params)
	})
	api.UserLogoutHandler = user.LogoutHandlerFunc(func(params user.LogoutParams) middleware.Responder {
		return server.Logout(params)
	})
	api.UserModifyUserPasswordHandler = user.ModifyUserPasswordHandlerFunc(func(params user.ModifyUserPasswordParams, principal *v1.Principal) middleware.Responder {
		return server.ModifyUserPassword(params, principal)
	})
	api.OauthTokenHandler = oauth.TokenHandlerFunc(func(params oauth.TokenParams) middleware.Responder {
		return server.Token(params)
	})
	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateUser(params, principal)
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
	server.RegisterClients()
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return mws.Limiter(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return mws.HandlePanic(mws.RedocUI(mws.LogViaLogrus(mws.Cross(handler))))
}

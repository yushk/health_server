// Code generated by go-swagger; DO NOT EDIT.

package oauth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TokenHandlerFunc turns a function with the right signature into a token handler
type TokenHandlerFunc func(TokenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TokenHandlerFunc) Handle(params TokenParams) middleware.Responder {
	return fn(params)
}

// TokenHandler interface for that can handle valid token params
type TokenHandler interface {
	Handle(TokenParams) middleware.Responder
}

// NewToken creates a new http.Handler for the token operation
func NewToken(ctx *middleware.Context, handler TokenHandler) *Token {
	return &Token{Context: ctx, Handler: handler}
}

/* Token swagger:route POST /v1/oauth/token oauth token

权限认证

权限认证

*/
type Token struct {
	Context *middleware.Context
	Handler TokenHandler
}

func (o *Token) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTokenParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

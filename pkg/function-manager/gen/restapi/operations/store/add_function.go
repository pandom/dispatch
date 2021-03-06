///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddFunctionHandlerFunc turns a function with the right signature into a add function handler
type AddFunctionHandlerFunc func(AddFunctionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddFunctionHandlerFunc) Handle(params AddFunctionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddFunctionHandler interface for that can handle valid add function params
type AddFunctionHandler interface {
	Handle(AddFunctionParams, interface{}) middleware.Responder
}

// NewAddFunction creates a new http.Handler for the add function operation
func NewAddFunction(ctx *middleware.Context, handler AddFunctionHandler) *AddFunction {
	return &AddFunction{Context: ctx, Handler: handler}
}

/*AddFunction swagger:route POST / Store addFunction

Add a new function

*/
type AddFunction struct {
	Context *middleware.Context
	Handler AddFunctionHandler
}

func (o *AddFunction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddFunctionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

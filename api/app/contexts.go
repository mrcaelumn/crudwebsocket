// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "crudws": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/mrcaelumn/crudwebsocket/api/design
// --out=$(GOPATH)/src/github.com/mrcaelumn/crudwebsocket/api
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// CreateCrudContext provides the crud create action context.
type CreateCrudContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCreateCrudContext parses the incoming request URL and body, performs validations and creates the
// context used by the crud controller create action.
func NewCreateCrudContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateCrudContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateCrudContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateCrudContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *CreateCrudContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateCrudContext) BadRequest(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateCrudContext) Unauthorized(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// Forbidden sends a HTTP response with status code 403.
func (ctx *CreateCrudContext) Forbidden(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 403, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateCrudContext) InternalServerError(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeleteCrudContext provides the crud delete action context.
type DeleteCrudContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID     string
	Idchat string
}

// NewDeleteCrudContext parses the incoming request URL and body, performs validations and creates the
// context used by the crud controller delete action.
func NewDeleteCrudContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteCrudContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteCrudContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	paramIdchat := req.Params["idchat"]
	if len(paramIdchat) > 0 {
		rawIdchat := paramIdchat[0]
		rctx.Idchat = rawIdchat
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteCrudContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteCrudContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteCrudContext) BadRequest(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *DeleteCrudContext) Unauthorized(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// Forbidden sends a HTTP response with status code 403.
func (ctx *DeleteCrudContext) Forbidden(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 403, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteCrudContext) InternalServerError(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// SelectCrudContext provides the crud select action context.
type SelectCrudContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewSelectCrudContext parses the incoming request URL and body, performs validations and creates the
// context used by the crud controller select action.
func NewSelectCrudContext(ctx context.Context, r *http.Request, service *goa.Service) (*SelectCrudContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SelectCrudContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SelectCrudContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *SelectCrudContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *SelectCrudContext) BadRequest(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SelectCrudContext) Unauthorized(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// Forbidden sends a HTTP response with status code 403.
func (ctx *SelectCrudContext) Forbidden(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 403, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *SelectCrudContext) InternalServerError(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// UpdateCrudContext provides the crud update action context.
type UpdateCrudContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewUpdateCrudContext parses the incoming request URL and body, performs validations and creates the
// context used by the crud controller update action.
func NewUpdateCrudContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateCrudContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateCrudContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateCrudContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateCrudContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateCrudContext) BadRequest(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *UpdateCrudContext) Unauthorized(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// Forbidden sends a HTTP response with status code 403.
func (ctx *UpdateCrudContext) Forbidden(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 403, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateCrudContext) InternalServerError(r *CrudwebsocketError) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// VersionVersionContext provides the version version action context.
type VersionVersionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewVersionVersionContext parses the incoming request URL and body, performs validations and creates the
// context used by the version controller version action.
func NewVersionVersionContext(ctx context.Context, r *http.Request, service *goa.Service) (*VersionVersionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := VersionVersionContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *VersionVersionContext) OK(r *CrudwebsocketVersion) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

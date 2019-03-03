// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "crudws": Application Controllers
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
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// CrudController is the controller interface for the Crud actions.
type CrudController interface {
	goa.Muxer
	Create(*CreateCrudContext) error
	Delete(*DeleteCrudContext) error
	Select(*SelectCrudContext) error
	Update(*UpdateCrudContext) error
}

// MountCrudController "mounts" a Crud resource controller on the given service.
func MountCrudController(service *goa.Service, ctrl CrudController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCrudContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/crudws/create", ctrl.MuxHandler("create", h, nil))
	service.LogInfo("mount", "ctrl", "Crud", "action", "Create", "route", "POST /crudws/create")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCrudContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("GET", "/crudws/delete/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Crud", "action", "Delete", "route", "GET /crudws/delete/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSelectCrudContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Select(rctx)
	}
	service.Mux.Handle("GET", "/crudws/select/:id", ctrl.MuxHandler("select", h, nil))
	service.LogInfo("mount", "ctrl", "Crud", "action", "Select", "route", "GET /crudws/select/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateCrudContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("POST", "/crudws/update/:id", ctrl.MuxHandler("update", h, nil))
	service.LogInfo("mount", "ctrl", "Crud", "action", "Update", "route", "POST /crudws/update/:id")
}

// PublicController is the controller interface for the Public actions.
type PublicController interface {
	goa.Muxer
	goa.FileServer
}

// MountPublicController "mounts" a Public resource controller on the given service.
func MountPublicController(service *goa.Service, ctrl PublicController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/css/*filepath", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/index", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/js/*filepath", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/css/*filepath", "public/css/")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/css/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/css/", "route", "GET /css/*filepath")

	h = ctrl.FileHandler("/index", "public/html/index.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/index", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/html/index.html", "route", "GET /index")

	h = ctrl.FileHandler("/js/*filepath", "public/js/")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/js/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/js/", "route", "GET /js/*filepath")

	h = ctrl.FileHandler("/css/", "public/css/index.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/css/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/css/index.html", "route", "GET /css/")

	h = ctrl.FileHandler("/js/", "public/js/index.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/js/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/js/index.html", "route", "GET /js/")
}

// handlePublicOrigin applies the CORS response headers corresponding to the origin.
func handlePublicOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "api/swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "api/swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// VersionController is the controller interface for the Version actions.
type VersionController interface {
	goa.Muxer
	Version(*VersionVersionContext) error
}

// MountVersionController "mounts" a Version resource controller on the given service.
func MountVersionController(service *goa.Service, ctrl VersionController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewVersionVersionContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Version(rctx)
	}
	service.Mux.Handle("GET", "/crudws/version", ctrl.MuxHandler("version", h, nil))
	service.LogInfo("mount", "ctrl", "Version", "action", "Version", "route", "GET /crudws/version")
}
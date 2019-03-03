package controller

import (
	"github.com/goadesign/goa"
	"github.com/mrcaelumn/crudwebsocket/api/app"
	"github.com/mrcaelumn/crudwebsocket/version"
)

// VersionController implements the version resource.
type VersionController struct {
	*goa.Controller
}

// NewVersionController creates a version controller.
func NewVersionController(service *goa.Service) *VersionController {
	return &VersionController{Controller: service.NewController("VersionController")}
}

// Version runs the version action.
func (c *VersionController) Version(ctx *app.VersionVersionContext) error {
	// VersionController_Version: start_implement

	// Put your logic here
	res := &app.CrudwebsocketVersion{
		Version: version.Version,
		Git:     &version.GitCommit,
	}
	return ctx.OK(res)
	// VersionController_Version: end_implement
}

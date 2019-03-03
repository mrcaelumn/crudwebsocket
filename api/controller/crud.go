package controller

import (
	"github.com/goadesign/goa"
	"github.com/mrcaelumn/crudwebsocket/api/app"
)

// CrudController implements the crud resource.
type CrudController struct {
	*goa.Controller
}

// NewCrudController creates a crud controller.
func NewCrudController(service *goa.Service) *CrudController {
	return &CrudController{Controller: service.NewController("CrudController")}
}

// Create runs the create action.
func (c *CrudController) Create(ctx *app.CreateCrudContext) error {
	// CrudController_Create: start_implement

	// Put your logic here

	return nil
	// CrudController_Create: end_implement
}

// Delete runs the delete action.
func (c *CrudController) Delete(ctx *app.DeleteCrudContext) error {
	// CrudController_Delete: start_implement

	// Put your logic here

	return nil
	// CrudController_Delete: end_implement
}

// Select runs the select action.
func (c *CrudController) Select(ctx *app.SelectCrudContext) error {
	// CrudController_Select: start_implement

	// Put your logic here

	return nil
	// CrudController_Select: end_implement
}

// Update runs the update action.
func (c *CrudController) Update(ctx *app.UpdateCrudContext) error {
	// CrudController_Update: start_implement

	// Put your logic here

	return nil
	// CrudController_Update: end_implement
}

package main

import (
	"database/sql"
	"micro/app"

	"github.com/goadesign/goa"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
	*sql.DB
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service, db *sql.DB) *OperandsController {
	return &OperandsController{
		Controller: service.NewController("OperandsController"),
		DB:         db,
	}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	// OperandsController_Add: start_implement

	// Put your logic here

	return ctx.OK([]byte("Done !"))
	// OperandsController_Add: end_implement
}

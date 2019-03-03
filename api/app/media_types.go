// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "crudws": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/mrcaelumn/crudwebsocket/api/design
// --out=$(GOPATH)/src/github.com/mrcaelumn/crudwebsocket/api
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// Crudwebsocket-Error media type (default view)
//
// Identifier: application/vnd.crudwebsocket-error+json; view=default
type CrudwebsocketError struct {
	// Message ID
	Code string `form:"code" json:"code" yaml:"code" xml:"code"`
	// Localized message
	Msg string `form:"msg" json:"msg" yaml:"msg" xml:"msg"`
}

// Validate validates the CrudwebsocketError media type instance.
func (mt *CrudwebsocketError) Validate() (err error) {
	if mt.Code == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "code"))
	}
	if mt.Msg == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "msg"))
	}
	return
}

// Crudwebsocket-Version media type (default view)
//
// Identifier: application/vnd.crudwebsocket-version+json; view=default
type CrudwebsocketVersion struct {
	// Git commit hash
	Git *string `form:"git,omitempty" json:"git,omitempty" yaml:"git,omitempty" xml:"git,omitempty"`
	// Application version
	Version string `form:"version" json:"version" yaml:"version" xml:"version"`
}

// Validate validates the CrudwebsocketVersion media type instance.
func (mt *CrudwebsocketVersion) Validate() (err error) {
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

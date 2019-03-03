package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// swagger
var _ = Resource("swagger", func() {
	NoSecurity()
	Origin("*", func() {
		Methods("GET")
	})
	Files("swagger.json", "api/swagger/swagger.json")
})

// version
var VersionType = Type("version", func() {
	Attribute("version", String, "Application version", func() {
		Example("1.0")
	})
	Attribute("git", String, "Git commit hash", func() {
		Example("000000")
	})

	Required("version")
})

var _ = Resource("version", func() {
	Action("version", func() {
		Routing(GET("version"))
		Response(OK, VersionMedia)
		Metadata("swagger:summary", "Return application's version and commit hash")
	})
})

var _ = Resource("crud", func() {
	Action("select", func() {
		Routing(GET("/select/:id"))
		Params(func() {
			Param("id", String, "id of items", func() {
				Example("146380193.16379966.1494326266")
			})
		})
		Response(OK, "application/json")
		Response(NoContent)
		Response(Unauthorized)
		Response(Forbidden)
		Response(BadRequest, CustomeErrorMedia)
		Response(InternalServerError, CustomeErrorMedia)
		Metadata("swagger:summary", "Get data")
	})
	Action("delete", func() {
		Routing(GET("/delete/:id"))
		Params(func() {
			Param("id", String, "id of items", func() {
				Example("146380193.16379966.1494326266")
			})
		})
		Response(OK, "application/json")
		Response(NoContent)
		Response(Unauthorized)
		Response(Forbidden)
		Response(BadRequest, CustomeErrorMedia)
		Response(InternalServerError, CustomeErrorMedia)
		Metadata("swagger:summary", "Get data")
	})
	Action("create", func() {
		Routing(POST("/create"))
		Response(OK, "application/json")
		Response(NoContent)
		Response(Unauthorized)
		Response(Forbidden)
		Response(BadRequest, CustomeErrorMedia)
		Response(InternalServerError, CustomeErrorMedia)
		Metadata("swagger:summary", "Get data")
	})
	Action("update", func() {
		Routing(POST("/update/:id"))
		Params(func() {
			Param("id", String, "id of items", func() {
				Example("146380193.16379966.1494326266")
			})
		})
		Response(OK, "application/json")
		Response(NoContent)
		Response(Unauthorized)
		Response(Forbidden)
		Response(BadRequest, CustomeErrorMedia)
		Response(InternalServerError, CustomeErrorMedia)
		Metadata("swagger:summary", "Get data")
	})
})

var _ = Resource("public", func() {

	Origin("*", func() { // CORS policy that applies to all actions and file servers
		Methods("GET") // of "public" resource
	})

	Files("/index", "public/html/index.html")
	Files("/js/*filepath", "public/js/")   // Serve all files under the public/js directory
	Files("/css/*filepath", "public/css/") // Serve all files under the public/js directory
})

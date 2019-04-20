package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("adder", func() {
	apidsl.Title("The adder API")
	apidsl.Description("A teser for go")
	apidsl.Host("localhost:3000")
	apidsl.Scheme("http")
})

var _ = apidsl.Resource("operands", func() {
	apidsl.Action("add", func() {
		apidsl.Routing(apidsl.GET("/"))
		apidsl.Description("add returns the sum of the left and right parameters in the response body")
		apidsl.Response(design.OK, "text/plain")
	})
})

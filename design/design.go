package design

import (
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("adder", func() {
	apidsl.Title("The adder API")
	apidsl.Description("A teser for go")
	apidsl.Host("localhost:3000")
	apidsl.Scheme("http")
})

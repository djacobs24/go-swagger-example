// Go Swagger Example API
//
// The purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// There are no TOS at this moment.
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     Contact: David Jacobs<dav.david.jacobs@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	// Imported so the swagger generator will hit the user package
	"fmt"

	"github.com/djacobs24/go-swagger-example/models"
	"github.com/djacobs24/go-swagger-example/user"
)

func main() {
	user.GetUser()
	user := models.User{}
	fmt.Print(user)
}

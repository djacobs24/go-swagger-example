//     Schemes: http
//     Host: localhost:8080
//     Version: 0.0.1
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
	"github.com/djacobs24/go-swagger-example/user"
)

func main() {
	user.GetUser()
}

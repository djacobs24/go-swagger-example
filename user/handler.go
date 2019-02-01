package user

// swagger:operation GET /user/{id} user GetUser
//
// Get one user by ID
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the user
//   required: true
//   type: string
// responses:
//   '200':
//     description: The user object
//     schema:
//       $ref: "#/definitions/user"
func GetUser() {

}

// swagger:operation POST /user user PostUser
//
// Create one user
// ---
// parameters:
// - name: user
//   in: body
//   description: The user object
//   required: true
//   schema:
//     $ref: "#/definitions/user"
// responses:
//   '200':
//     description: The user object
//     schema:
//       $ref: "#/definitions/user"
func PostUser() {

}

// swagger:operation PUT /user/{id} user PutUser
//
// Update one user by ID.
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the user
//   required: true
//   type: string
// responses:
//   '200':
//     description: The user object
//     schema:
//       $ref: "#/definitions/user"
func PutUser() {

}

// swagger:operation DELETE /user/{id} user DeleteUser
//
// Delete one user by ID.
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the user
//   required: true
//   type: string
// responses:
//   '200':
//     description: The user object
//     schema:
//       $ref: "#/definitions/user"
func DeleteUser() {

}

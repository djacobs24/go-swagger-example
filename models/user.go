package models

// User user
// swagger:model user
type User struct {
	// name
	Name string `json:"name,omitempty"`
	// username
	Username string `json:"username"`
	// job_title
	JobTitle string `json:"job_title"`
	// location
	Location string `json:"location"`
}

package dto

type CreateUserInput struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	RoleId   int    `json:"role_id"`
}

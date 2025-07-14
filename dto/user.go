package dto

type CreateUserRequest struct {
	Fullname *string
	Username string
	Email    string
	Password string
}

type UpdateUserRequest struct {
	Id       int64
	Fullname string
	Username string
	Email    string
}

package dto

import "time"

// Request
type BrowseUserRequest struct {
	Pagination
	Fullname *string `json:"fullname"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

type CreateUserRequest struct {
	Fullname  string `json:"fullname" validate:"max=100"`
	Username  string `json:"username" validate:"max=15"`
	Email     string `json:"email" validate:"max=150,email"`
	Password  string `json:"password"`
	CreatedBy string `json:"createdBy" swaggerignore:"true"`
}

type UpdateUserRequest struct {
	Id         int64  `json:"id" swaggerignore:"true"`
	Fullname   string `json:"fullname" validate:"max=100"`
	Username   string `json:"username" validate:"max=15"`
	Email      string `json:"email" validate:"max=150,email"`
	ModifiedBy string `json:"modifiedBy" swaggerignore:"true"`
}

type UpdateUserPasswordRequest struct {
	UserId      int64  `json:"userId" validate:"required" swaggerignore:"true"`
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
	ModifiedBy  string `json:"modifiedBy" swaggerignore:"true"`
}

type DeleteUserRequest struct {
	Id        int64
	DeletedBy string
}

//Response
type UserResponse struct {
	Id         int64      `json:"id"`
	Fullname   string     `json:"fullname"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Active     bool       `json:"active"`
	CreatedBy  string     `json:"createdBy"`
	CreatedAt  time.Time  `json:"createdAt"`
	ModifiedBy *string    `json:"modifiedBy"`
	ModifiedAt *time.Time `json:"modifiedAt"`
}

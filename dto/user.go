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
	Fullname  string
	Username  string
	Email     string
	Password  string
	CreatedBy string
}

type UpdateUserRequest struct {
	Id         int64
	Fullname   string
	Username   string
	Email      string
	ModifiedBy string
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

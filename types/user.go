package types

type User struct {
	Email string `json:"email"`
	Age   int64  `json:"age"`
}
type UserResponse struct {
	*ApiResponse
	*User
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type UpdateRequest struct {
	Email     string `json:"email" binding:"required"`
	UpdateAge int64  `json:"updateAge" binding:"required"`
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteRequest struct {
	Email string `json:"email" binding:"required"`
}

func (c *DeleteRequest) ToUser() *User {
	return &User{
		Email: c.Email,
	}
}

type DeleteUserResponse struct {
	*ApiResponse
}

type ErrorResponse struct {
	*ApiResponse
}

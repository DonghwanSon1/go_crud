package types

type User struct {
	Email string `json:"email"`
	Age   int64  `json:"age"`
}
type UserResponse struct {
	*ApiResponse
	*User
}

//type SignupRequest struct {
//	Email    string `json:"email" binding:"required"`
//	Password string `json:"password" binding:"required"`
//	Age      int64  `json:"age" binding:"required"`
//}
//type SignupUserResponse struct {
//	*ApiResponse
//}

//func (c *SignupRequest) SignupRq() *models.User {
//	return &models.User{
//		Email:    c.Email,
//		Password: c.Password,
//		Age:      c.Age,
//	}
//}

//type LoginRequest struct {
//	Email    string `json:"email" binding:"required"`
//	Password string `json:"password" binding:"required"`
//}
//
//type LoginResponse struct {
//	*ApiResponse
//	Token string `json:"token"`
//}

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

package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
type UserResponse struct {
	*ApiResponse
	*User
}

type CreateUserResponse struct {
	*ApiResponse
}

type GetUserResponse struct {
	*ApiResponse
	*User
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserResponse struct {
	*ApiResponse
}

package usersInfo

import (
	"go_crud/models"
	"go_crud/types"
	"gorm.io/gorm"
	"time"
)

type Response struct {
	*types.ApiResponse
	//*models.UsersInfo
}

type SignupRequest struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	About    string `json:"about"`
	Gender   string `json:"gender" binding:"required"`
	Birth    string `json:"birth" binding:"required,datetime=2006-01-02"`
}

func (rq *SignupRequest) CreateUsersInfo() *models.UsersInfo {
	birth, _ := time.Parse("2006-01-02", rq.Birth)
	return &models.UsersInfo{
		UserId:    rq.UserId,
		Password:  rq.Password,
		Name:      rq.Name,
		Phone:     rq.Phone,
		Email:     rq.Email,
		About:     rq.About,
		Gender:    rq.Gender,
		Birth:     birth,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}
}

type SignupUserResponse struct {
	*types.ApiResponse
}

type LoginRequest struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	*types.ApiResponse
	Token string `json:"token"`
}

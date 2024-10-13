package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/internal/domain"
	"github.com/kynmh69/mormorare/pkg/logging"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type User struct {
	UserName string    `form:"username" binding:"required"`
	Password string    `form:"password" binding:"required"`
	Email    string    `form:"email" binding:"required"`
	BirthDay time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"0" binding:"required"`
}

type DeleteUser struct {
	UserName string `uri:"username" binding:"required" example:"test1"`
}

type UserResponse struct {
	UserName string    `json:"username" example:"test1"`
	Email    string    `json:"email" example:"example@test.com"`
	BirthDay time.Time `json:"birth_day" example:"2006-01-02"`
}

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

// Create
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user_name body string true "set username"
// @Param password body string true "set password"
// @Param email body string true "set email"
// @Param birthday body string true "set birthday"
// @Success 201 {object} UserResponse
// @Failure 400 {object} domain.ErrorJson
// @Router /api/v1/users [post]
func (u *UserHandler) Create(ctx *gin.Context) {
	logger := logging.GetLogger()
	// Create user
	var newUser User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	dUser := domain.NewUser(
		newUser.UserName,
		newUser.Password,
		newUser.Email,
		newUser.BirthDay,
	)
	u.db.Create(dUser)
	resp := UserResponse{
		UserName: dUser.UserName,
		Email:    dUser.Email,
		BirthDay: dUser.Birthday,
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (u *UserHandler) Update(ctx *gin.Context) {
	// Update user
	ctx.Status(http.StatusNotImplemented)
}

// Retrieve
// @Summary Retrieve all users
// @Description Retrieve all users
// @Tags users
// @Produce json
// @Success 200 {array} []UserResponse
// @Failure 400 {object} domain.ErrorJson
// @Router /api/v1/users [get]
func (u *UserHandler) Retrieve(ctx *gin.Context) {
	// Retrieve user
	var users []domain.User
	if err := u.db.Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	resp := make([]UserResponse, 0, len(users))
	for _, user := range users {
		resp = append(resp, UserResponse{
			UserName: user.UserName,
			Email:    user.Email,
			BirthDay: user.Birthday,
		})
	}
	if len(resp) == 0 {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// Delete
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Param user_name path string true "set username"
// @Success 204
// @Failure 400 {object} domain.ErrorJson
// @Router /api/v1/users/{user_name} [delete]
func (u *UserHandler) Delete(ctx *gin.Context) {
	// Delete user
	var deleteUser DeleteUser
	if err := ctx.ShouldBindUri(&deleteUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	if err := u.db.Where("user_name = ?", deleteUser.UserName).Delete(&domain.User{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	ctx.Status(http.StatusNoContent)
}

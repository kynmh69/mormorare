package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/internal/domain"
	"github.com/kynmh69/mormorare/internal/domain/repository"
	"github.com/kynmh69/mormorare/pkg/hash"
	"net/http"
	"time"
)

type User struct {
	UserName string    `form:"username" binding:"required"`
	Password string    `form:"password" binding:"required,min=8,max=32"`
	Email    string    `form:"email" binding:"required,email"`
	BirthDay time.Time `form:"birthday" binding:"required" time_format:"2006-01-02"`
}

type UserId struct {
	Username string `uri:"id" binding:"required"`
}

type UserUpdate struct {
	UserName string `from:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

type DeleteUser struct {
	UserName string `uri:"username" binding:"required"`
}

type UserResponse struct {
	UserName string    `json:"username" example:"test1"`
	Email    string    `json:"email" example:"example@test.com"`
	BirthDay time.Time `json:"birthday" example:"2006-01-02"`
}

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// Create
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param username body string true "set username"
// @Param password body string true "set password"
// @Param email body string true "set email"
// @Param birthday body string true "set birthday"
// @Success 201 {object} UserResponse
// @Failure 400 {object} domain.ErrorJson
// @Router /api/v1/users [post]
func (u *UserHandler) Create(ctx *gin.Context) {
	// Create user
	var newUser User
	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}

	hashedPassword, err := hash.HashPassword(newUser.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	dUser := domain.NewUser(
		newUser.UserName,
		hashedPassword,
		newUser.Email,
		newUser.BirthDay,
	)
	if err := u.repo.CreateUser(dUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	resp := UserResponse{
		UserName: dUser.UserName,
		Email:    dUser.Email,
		BirthDay: dUser.Birthday,
	}
	ctx.JSON(http.StatusCreated, resp)
}

// Update
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param username body string false "set username"
// @Param password body string false "set password"
// @Param email body string false "set email"
// @Success 201 {object} UserResponse
// @Failure 400 {object} domain.ErrorJson
// @Router /api/v1/users/{username} [put]
func (u *UserHandler) Update(ctx *gin.Context) {
	var (
		updateUser UserUpdate
		dUser      *domain.User
		uriUser    UserId
		err        error
	)
	if err = ctx.ShouldBindUri(&uriUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	if err = ctx.ShouldBind(&updateUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	dUser, err = u.repo.GetUserByUsername(uriUser.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	if updateUser.Password != "" {
		hashedPassword, err := hash.HashPassword(updateUser.Password)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
			return
		}
		dUser.Password = hashedPassword
	}
	if updateUser.Email != "" {
		dUser.Email = updateUser.Email
	}
	if updateUser.UserName != "" {
		dUser.UserName = updateUser.UserName
	}
	if err := u.repo.UpdateUser(dUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	// Update user
	resp := UserResponse{
		UserName: dUser.UserName,
		Email:    dUser.Email,
		BirthDay: dUser.Birthday,
	}
	ctx.JSON(http.StatusCreated, resp)
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
	var (
		users []domain.User
		err   error
	)
	users, err = u.repo.GetUsers()
	if err != nil {
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
// @Router /api/v1/users/{username} [delete]
func (u *UserHandler) Delete(ctx *gin.Context) {
	// Delete user
	var deleteUser DeleteUser
	if err := ctx.ShouldBindUri(&deleteUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	user, err := u.repo.GetUserByUsername(deleteUser.UserName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	if err := u.repo.DeleteUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	ctx.Status(http.StatusNoContent)
}

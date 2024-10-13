package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/internal/domain"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type User struct {
	UserName string    `binding:"required"`
	Password string    `binding:"required"`
	Email    string    `binding:"required"`
	BirthDay time.Time `binding:"required" time_format:"2006-01-02"`
}

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (u *UserHandler) Create(ctx *gin.Context) {
	// Create user
	var newUser User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
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
	ctx.JSON(http.StatusCreated, dUser)
}

func (u *UserHandler) Update(ctx *gin.Context) {
	// Update user
	ctx.Status(http.StatusNotImplemented)
}

func (u *UserHandler) Retrieve(ctx *gin.Context) {
	// Retrieve user
	var users []domain.User
	if err := u.db.Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorJson(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserHandler) Delete(ctx *gin.Context) {
	// Delete user
	ctx.Status(http.StatusNotImplemented)
}

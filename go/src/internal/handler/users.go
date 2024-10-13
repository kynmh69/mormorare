package handler

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u UserHandler) Create() {
	// Create user
}

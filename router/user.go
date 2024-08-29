package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"personal_blog/entity"
	"personal_blog/logger"
	"personal_blog/service"
)

type UserRouter struct {
	userService *service.UserService
	logger      logger.Logger
}

func NewUserRouter(userService *service.UserService, logger logger.Logger) *UserRouter {
	return &UserRouter{
		userService: userService,
		logger:      logger,
	}
}

func (router *UserRouter) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	uintID, err := convertStringToUint64(id)
	if err != nil {
		router.logger.Error("user convert string to uint64 error", err)
		return InternalServerError("user")
	}
	user, err := router.userService.FindByID(uintID)
	if err != nil {
		router.logger.Error("can't find the user", err)
		return NotFoundError("user")
	}
	return ctx.JSON(http.StatusOK, user)
}
func (router *UserRouter) Create(ctx echo.Context) error {
	var user entity.User
	if err := ctx.Bind(&user); err != nil {
		router.logger.Error("create user bind data error", err)
		return BadRequestError("invalid user data")
	}
	err := router.userService.CreateUser(&user)
	if err != nil {
		router.logger.Error("create user error", err)
		return InternalServerError("create user has an internal error")
	}
	return ctx.JSON(http.StatusOK, user)
}
func (router *UserRouter) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	var user entity.User
	if err := ctx.Bind(&user); err != nil {
		router.logger.Error("update user bind data error", err)
		return BadRequestError("update user has an internal error")
	}
	uintID, err := convertStringToUint64(id)
	if err != nil {
		router.logger.Error("user convert string to uint64 error", err)
		return InternalServerError("update user has an internal error")
	}
	user.ID = uintID
	if err := router.userService.UpdateUser(&user); err != nil {
		router.logger.Error("update user error", err)
		return BadRequestError("invalid user data")
	}
	return ctx.JSON(http.StatusOK, user)
}
func (router *UserRouter) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	uintID, err := convertStringToUint64(id)
	if err != nil {
		router.logger.Error("convert string to uint64 error", err)
		return InternalServerError("delete user has an internal error")
	}
	if err := router.userService.DeleteUser(uintID); err != nil {
		router.logger.Error("delete user error", err)
		return InternalServerError("delete user has an internal error")
	}
	return ctx.NoContent(http.StatusNoContent)
}

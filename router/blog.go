package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"personal_blog/entity"
	"personal_blog/log"
	"personal_blog/service"
)

type BlogRouter struct {
	blogService *service.BlogService
	logger      log.Logger
}

func NewBlogRouter(blogService *service.BlogService, logger log.Logger) *BlogRouter {
	return &BlogRouter{
		blogService: blogService,
		logger:      logger,
	}
}

func (router *BlogRouter) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	uintID, err := convertStringToUint64(id)
	if err != nil {
		router.logger.Error("blog convert string to uint64 error", err)
		return InternalServerError("blog")
	}
	blog, err := router.blogService.FindByID(uintID)
	if err != nil {
		router.logger.Error("can't find blog", err)
		return NotFoundError("blog not found")
	}
	return ctx.JSON(http.StatusOK, blog)
}
func (router *BlogRouter) Create(ctx echo.Context) error {
	var blog entity.Blog
	if err := ctx.Bind(&blog); err != nil {
		router.logger.Error("bind blog error", err)
		return BadRequestError("invalid blog data")
	}
	err := router.blogService.CreateBlog(&blog)
	if err != nil {
		router.logger.Error("create new blog error", err)
		return InternalServerError("create blog has an internal error")
	}
	return ctx.JSON(http.StatusCreated, blog)
}
func (router *BlogRouter) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	var blog entity.Blog
	if err := ctx.Bind(&blog); err != nil {
		router.logger.Error("update blog bind data error", err)
		return BadRequestError("update blog has an internal error")
	}
	uintID, err := convertStringToUint64(id)
	if err != nil {
		router.logger.Error("convert string to uint64 error", err)
		return InternalServerError("update blog has an internal error")
	}
	blog.ID = uintID
	err = router.blogService.UpdateBlog(&blog)
	if err != nil {
		router.logger.Error("update blog error", err)
		return BadRequestError("invalid blog data")
	}
	return ctx.JSON(http.StatusOK, blog)
}
func (router *BlogRouter) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	uintID, err := convertStringToUint64(id)
	if err != nil {
		router.logger.Error("convert string to uint64 error", err)
		return InternalServerError("delete blog has an internal error")
	}
	err = router.blogService.DeleteBlog(uintID)
	if err != nil {
		router.logger.Error("delete blog error", err)
		return InternalServerError("delete blog has an internal error")
	}
	return ctx.NoContent(http.StatusNoContent)
}

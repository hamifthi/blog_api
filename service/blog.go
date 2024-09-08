package service

import (
	"errors"
	"personal_blog/entity"
	"personal_blog/log"
	"personal_blog/repository"
)

type BlogService struct {
	blogRepository repository.Blog
	logger         log.Logger
}

func NewBlogService(blogRepository repository.Blog, logger log.Logger) *BlogService {
	return &BlogService{blogRepository: blogRepository, logger: logger}
}

func (bs *BlogService) FindByID(id uint) (*entity.Blog, error) {
	blog, err := bs.blogRepository.GetByID(id)
	if err != nil {
		bs.logger.Error("facing error while finding blog by ID: ", err, id)
		return nil, err
	}
	if blog == nil {
		bs.logger.Warn("Couldn't find blog by ID: ", id)
		return nil, errors.New("blog not found")
	}
	return blog, nil
}

func (bs *BlogService) Create(blog *entity.Blog) error {
	if err := validateBlog(blog.Title, blog.Content, blog.AuthorID); err != nil {
		bs.logger.Error("error in blog validation: ", err, blog)
		return err
	}
	_, err := bs.blogRepository.Create(blog)
	if err != nil {
		bs.logger.Error("facing error while creating blog by title, content, authorID: ",
			err, blog)
		return err
	}
	return nil
}

func (bs *BlogService) Update(blog *entity.Blog) error {
	// Validate input
	if err := validateBlog(blog.Title, blog.Content, blog.AuthorID); err != nil {
		bs.logger.Error("facing error while validating blog by ID: ", err, blog)
		return err
	}

	_, err := bs.FindByID(blog.ID)
	if err != nil {
		bs.logger.Error("facing error while finding blog by ID: ", err, blog.ID)
		return err
	}

	err = bs.blogRepository.Update(blog)
	if err != nil {
		bs.logger.Error("facing error while updating blog by ID: ", err, blog)
		return err
	}
	return nil
}

func (bs *BlogService) Delete(id uint) error {
	blog, err := bs.FindByID(id)
	if err != nil {
		bs.logger.Error("facing error while deleting blog by ID: ", err, id)
		return err
	}
	return bs.blogRepository.Delete(blog)
}

package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"os"
	"personal_blog/conf"
	"personal_blog/log"
	"personal_blog/repository"
	"personal_blog/router"
	"personal_blog/service"
	"personal_blog/storage"
)

func main() {
	// initialize logger
	developmentConf := log.DevelopmentConfig()
	logger, err := log.NewLogger(developmentConf)
	if err != nil {
		panic(err)
	}

	// load config
	config, err := conf.LoadConfig("resources", "app", "yml")
	if err != nil {
		logger.Error("load config err", zap.Error(err))
		os.Exit(1)
	}
	// create db connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Database.Host, config.Database.User, config.Database.Password,
		config.Database.DBName, config.Database.Port, config.Database.SSLMode,
	)
	dbGenerator := &storage.DBGenerator{}
	if err := dbGenerator.ConnectDB(postgres.Open(dsn)); err != nil {
		os.Exit(1)
	}

	// migrate the db
	if err := dbGenerator.MigrateDB(); err != nil {
		logger.Error("migration err", zap.Error(err))
		os.Exit(1)
	}

	// create repositories here
	blogRepository := repository.NewBlogRepository(dbGenerator.DB)
	userRepository := repository.NewUserRepository(dbGenerator.DB)

	// create services
	blogService := service.NewBlogService(blogRepository, logger)
	userService := service.NewUserService(userRepository, logger)

	// create router
	blogRouter := router.NewBlogRouter(blogService, logger)
	userRouter := router.NewUserRouter(userService, logger)

	// Create a new Echo instance
	echoRouter := echo.New()
	echoRouter.POST("/users", userRouter.Create)
	echoRouter.GET("/users/:id", userRouter.Get)
	echoRouter.PUT("/users/:id", userRouter.Update)
	echoRouter.DELETE("/users/:id", userRouter.Delete)

	echoRouter.POST("/blogs", blogRouter.Create)
	echoRouter.GET("/blogs/:id", blogRouter.Get)
	echoRouter.PUT("/blogs/:id", blogRouter.Update)
	echoRouter.DELETE("/blogs/:id", blogRouter.Delete)

	echoRouter.Logger.Fatal(echoRouter.Start(":8080"))
}

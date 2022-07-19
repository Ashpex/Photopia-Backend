package main

import (
	"example.com/gallery/config"
	"example.com/gallery/controller"
	"example.com/gallery/helper"
	"example.com/gallery/middleware"
	"example.com/gallery/repository"
	"example.com/gallery/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

var (
	// db is a global variable that represents the config connection
	db *gorm.DB = config.SetupDB()
	// Database repository
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	postRepository  repository.PostRepository  = repository.NewPostRepository(db)
	topicRepository repository.TopicRepository = repository.NewTopicRepository(db)
	// jwtService is a global variable that represents the jwt service (json web token)
	jwtService helper.JWTService = helper.NewJWTService()
	// Authentication service and controller
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	// User service and controller
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService, jwtService)

	// Post service and controller
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService, jwtService)
	// Topic service and controller
	topicService    service.TopicService       = service.NewTopicService(topicRepository)
	topicController controller.TopicController = controller.NewTopicController(topicService, jwtService)
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	defer config.CloseDB(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	postRoutes := r.Group("api/posts", middleware.AuthorizeJWT(jwtService))
	{
		postRoutes.GET("/", postController.All)
		postRoutes.POST("/", postController.Insert)
		postRoutes.GET("/:id", postController.FindByID)
		postRoutes.PUT("/:id", postController.Update)
		postRoutes.DELETE("/:id", postController.Delete)
		postRoutes.GET("/topic/:id", postController.FindByTopicID)
	}
	topicRoutes := r.Group("api/topics")
	{
		topicRoutes.GET("/", topicController.All)
		topicRoutes.POST("/", topicController.Insert)
		topicRoutes.GET("/:id", topicController.FindByID)
	}

	err := r.Run()
	if err != nil {
		return
	}
}

package controller

import (
	"example.com/gallery/dto"
	"example.com/gallery/entity"
	"example.com/gallery/helper"
	"example.com/gallery/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
)

// PostController is a contract about something that this controller can do
type PostController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

// postController is a struct to implement PostController
type postController struct {
	postService service.PostService
	jwtService  service.JWTService
}

// NewPostController is a function to create a new instance of PostController
func NewPostController(postServ service.PostService, jwtServ service.JWTService) PostController {
	return &postController{
		postService: postServ,
		jwtService:  jwtServ,
	}
}

// All is a function to get all posts
func (c *postController) All(context *gin.Context) {
	var posts []entity.Post = c.postService.All()
	response := helper.BuildResponse(true, "Get all posts successfully", posts)
	context.JSON(http.StatusOK, response)
}

// FindByID is a function to find a post by id
func (c *postController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var post entity.Post = c.postService.FindPostByID(id)
	if (post == entity.Post{}) {
		response := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
	} else {
		response := helper.BuildResponse(true, "Get post successfully", post)
		context.JSON(http.StatusOK, response)
	}
}

func (c *postController) Insert(context *gin.Context) {
	var postCreateDTO dto.PostCreateDTO
	errDTO := context.BindJSON(&postCreateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Invalid request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)

	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserId, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			postCreateDTO.UserID = convertedUserId
		}
		result := c.postService.Insert(postCreateDTO)
		response := helper.BuildResponse(true, "Insert post successfully", result)
		context.JSON(http.StatusOK, response)

	}
}

func (c *postController) Update(context *gin.Context) {
	var postUpdateDTO dto.PostUpdateDTO
	errDTO := context.ShouldBind(&postUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.postService.IsAllowedToEdit(userID, postUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			postUpdateDTO.UserID = id
		}
		result := c.postService.Update(postUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *postController) Delete(context *gin.Context) {
	var post entity.Post
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get ID", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	post.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.postService.IsAllowedToEdit(userID, post.ID) {
		c.postService.Delete(post)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
func (c *postController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

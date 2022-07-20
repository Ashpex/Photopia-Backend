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

type CommentController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	FindByPostID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type commentController struct {
	commentService service.CommentService
	jwtService     helper.JWTService
}

func NewCommentController(commentService service.CommentService, jwtService helper.JWTService) CommentController {
	return &commentController{
		commentService: commentService,
		jwtService:     jwtService,
	}
}

func (controller *commentController) All(context *gin.Context) {
	comments := controller.commentService.All()
	response := helper.BuildResponse(true, "Get all comments successfully", comments)
	context.JSON(http.StatusOK, response)
}

func (controller *commentController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var comment entity.Comment = controller.commentService.FindByID(id)
	if (comment == entity.Comment{}) {
		response := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
	} else {
		response := helper.BuildResponse(true, "Found comment", comment)
		context.JSON(http.StatusOK, response)
	}

}

func (controller *commentController) FindByPostID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var comments []entity.Comment = controller.commentService.FindByPostID(id)
	if comments == nil {
		response := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
	} else {
		response := helper.BuildResponse(true, "Found comments", comments)
		context.JSON(http.StatusOK, response)
	}

}

func (c *commentController) Insert(context *gin.Context) {
	var commentCreateDTO dto.CommentCreateDTO
	err := context.ShouldBind(&commentCreateDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			commentCreateDTO.UserID = convertedUserID
		}
		result := c.commentService.Insert(commentCreateDTO)
		response := helper.BuildResponse(true, "Insert comment successfully", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (controller *commentController) Update(context *gin.Context) {
	var commentUpdateDTO dto.CommentUpdateDTO
	err := context.ShouldBind(&commentUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, err := controller.jwtService.ValidateToken(authHeader)
	if err != nil {
		fmt.Sprintf("%v", err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if controller.commentService.IsAllowedToEdit(userID, commentUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			commentUpdateDTO.UserID = id
		}
		result := controller.commentService.Update(commentUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (controller *commentController) Delete(context *gin.Context) {
	var comment entity.Comment
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get the id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	comment.ID = id
	authHeader := context.GetHeader("Authorization")
	token, err := controller.jwtService.ValidateToken(authHeader)
	if err != nil {
		fmt.Sprintf("%v", err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if controller.commentService.IsAllowedToEdit(userID, comment.ID) {
		controller.commentService.Delete(comment)
		response := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (context *commentController) getUserIDByToken(token string) string {
	aToken, err := context.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

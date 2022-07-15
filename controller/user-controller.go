package controller

import (
	"example.com/gallery/dto"
	"example.com/gallery/helper"
	"example.com/gallery/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
)

// UserController is a contract about something that this controller can do
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

// NewUserController is a function to create a new instance of userController
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (uc *userController) Update(context *gin.Context) {
	userUpdateDTO := dto.UserUpdateDTO{}
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := uc.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Failed to process request", errToken.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		panic(errToken.Error())
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID, errID := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if errID != nil {
		response := helper.BuildErrorResponse("Failed to process request", errID.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		panic(errID.Error())
		return
	}
	userUpdateDTO.ID = userID

	u := uc.userService.Update(userUpdateDTO)
	response := helper.BuildResponse(true, "Successfully updated user", u)
	context.JSON(http.StatusOK, response)
}

func (uc *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := uc.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	user := uc.userService.Profile(userID)
	response := helper.BuildResponse(true, "Successfully retrieved user", user)
	context.JSON(http.StatusOK, response)
}

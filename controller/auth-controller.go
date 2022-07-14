package controller

import (
	"example.com/gallery/dto"
	"example.com/gallery/entity"
	"example.com/gallery/helper"
	"example.com/gallery/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (ac *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := ac.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken, _ := ac.jwtService.GenerateToken(v.ID)
		v.Token = generatedToken
		response := helper.BuildResponse(true, "Login successful", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("Failed to process request", "Invalid username or password", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}

func (ac *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !ac.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Email already exists", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := ac.authService.CreateUser(registerDTO)
		token, _ := ac.jwtService.GenerateToken(createdUser.ID)
		createdUser.Token = token
		response := helper.BuildResponse(true, "Registration successful", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

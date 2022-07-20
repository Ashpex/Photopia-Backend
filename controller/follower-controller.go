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

type FollowerController interface {
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	AllFollowers(c *gin.Context)
	AllFollowing(c *gin.Context)
}

type followerController struct {
	followerService service.FollowService
	jwtService      helper.JWTService
}

//NewPostController create a new instances of PostController
func NewFollowerController(followerService service.FollowService, jwtServ helper.JWTService) FollowerController {
	return &followerController{
		followerService: followerService,
		jwtService:      jwtServ,
	}
}

//All is a function that get all followers of a user
func (c *followerController) AllFollowers(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("user_id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param user_id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var followers []entity.Follower = c.followerService.AllFollowers(userId)
	response := helper.BuildResponse(true, "Get all followers successfully", followers)
	context.JSON(http.StatusOK, response)
}

//AllFollowing is a function that get all following of a user
func (c *followerController) AllFollowing(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("user_id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param user_id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var following []entity.Follower = c.followerService.AllFollowing(userId)
	response := helper.BuildResponse(true, "Get all following successfully", following)
	context.JSON(http.StatusOK, response)
}

func (c *followerController) Follow(context *gin.Context) {
	var followerFollowDTO dto.FollowDTO
	err := context.BindJSON(&followerFollowDTO)
	if err != nil {
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			followerFollowDTO.TargetUserID = convertedUserID
		}
		result := c.followerService.Follow(followerFollowDTO)
		response := helper.BuildResponse(true, "Follow sucessfully", result)
		context.JSON(http.StatusCreated, response)
	}

}

func (c *followerController) Unfollow(context *gin.Context) {
	var follower entity.Follower
	err := context.BindJSON(&follower)
	if err != nil {
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	c.followerService.UnFollow(follower)
	res := helper.BuildResponse(true, "Unfollow successfully", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}

func (c *followerController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		fmt.Println("Can not get user id from token: ", err)
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

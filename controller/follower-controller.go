package controller

import (
	"example.com/gallery/dto"
	"example.com/gallery/entity"
	"example.com/gallery/helper"
	"example.com/gallery/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FollowerController interface {
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	All(c *gin.Context)
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

func (c *followerController) All(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("user_id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var followers []entity.Follower = c.followerService.AllFollowers(id)
	response := helper.BuildResponse(true, "Get all followers successfully", followers)
	context.JSON(http.StatusOK, response)
}

func (c *followerController) Follow(context *gin.Context) {
	var followerFollowDTO dto.FollowDTO
	err := context.BindJSON(&followerFollowDTO)
	if err != nil {
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	follower := c.followerService.Follow(followerFollowDTO)
	response := helper.BuildResponse(true, "Follow successfully", follower)
	context.JSON(http.StatusOK, response)
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

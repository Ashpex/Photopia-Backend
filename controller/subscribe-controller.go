package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/dto"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/entity"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/helper"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/service"
	"net/http"
	"strconv"
)

type SubscribeController interface {
	Subscribe(context *gin.Context)
	Unsubscribe(context *gin.Context)
	AllSubscribes(context *gin.Context)
	CountSubscribes(context *gin.Context)
}

type subscribeController struct {
	subscribeService service.SubscribeService
	topicService     service.TopicService
	jwtService       helper.JWTService
}

func NewSubscribeController(subscribeService service.SubscribeService, jwtService helper.JWTService, topicService service.TopicService) SubscribeController {
	return &subscribeController{
		subscribeService: subscribeService,
		jwtService:       jwtService,
		topicService:     topicService,
	}
}

func (c *subscribeController) Subscribe(context *gin.Context) {
	var susbcribeDTO dto.SubscribeDTO

	topicId, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param topic_id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var topic entity.Topic = c.topicService.FindByID(topicId)
	if (topic == entity.Topic{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		token, err := c.jwtService.ValidateToken(authHeader)
		if err != nil {
			fmt.Sprintf("%v", err.Error())
		}
		claims := token.Claims.(jwt.MapClaims)
		userID := fmt.Sprintf("%v", claims["user_id"])
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			susbcribeDTO.UserID = id
		}
		susbcribeDTO.TopicID = topicId
		result := c.subscribeService.Subscribe(susbcribeDTO)
		res := helper.BuildResponse(true, "Subscribe successfully", result)
		context.JSON(http.StatusOK, res)
	}
}

func (c *subscribeController) Unsubscribe(context *gin.Context) {
	var subscribe entity.Subscribe
	topicId, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get the id", "No param topic_id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
		return
	}
	var topic entity.Topic = c.topicService.FindByID(topicId)
	if (topic == entity.Topic{}) {
		res := helper.BuildErrorResponse("Data not found", "No topic with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
		return
	} else {
		subscribe.TopicID = topicId
		authHeader := context.GetHeader("Authorization")
		token, err := c.jwtService.ValidateToken(authHeader)
		if err != nil {
			fmt.Sprintf("%v", err.Error())
		}
		claims := token.Claims.(jwt.MapClaims)
		userID := fmt.Sprintf("%v", claims["user_id"])
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			subscribe.UserID = id
			c.subscribeService.Unsubscribe(subscribe)
			res := helper.BuildResponse(true, "Unsubscribed successfully", helper.EmptyObj{})
			context.JSON(http.StatusOK, res)
		} else {
			res := helper.BuildErrorResponse("Failed to get the id", "No param user_id were found", helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)
		}
	}

}

func (c *subscribeController) AllSubscribes(context *gin.Context) {
	var subscribes []entity.Subscribe
	topicId, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get the id", "No param topic_id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	var topic entity.Topic = c.topicService.FindByID(topicId)
	if (topic == entity.Topic{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		subscribes = c.subscribeService.AllSubscribes(topicId)
		res := helper.BuildResponse(true, "All subscribes retrieved successfully", subscribes)
		context.JSON(http.StatusOK, res)
	}

}

func (c *subscribeController) CountSubscribes(context *gin.Context) {
	var numberOfSubscribes int
	topicId, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get the id", "No param topic_id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	var topic entity.Topic = c.topicService.FindByID(topicId)
	if (topic == entity.Topic{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		numberOfSubscribes = c.subscribeService.CountSubscribes(topicId)
		res := helper.BuildResponse(true, "Get number of subscribers successfully", numberOfSubscribes)
		context.JSON(http.StatusOK, res)
	}

}

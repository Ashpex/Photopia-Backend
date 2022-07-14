package middleware

import (
	"example.com/gallery/helper"
	"example.com/gallery/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

// AuthorizeJWT validates the token, returns error if token is invalid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token provided", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("User ID: ", claims["user_id"])
			log.Println("issuer: ", claims["issuer"])
		} else {
			log.Println("Token is invalid: ", err)
			response := helper.BuildErrorResponse("Failed to process request", "Invalid token", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}

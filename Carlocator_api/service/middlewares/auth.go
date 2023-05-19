package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/models"
	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "you need to login first",
			})
			c.Abort()
			return
		}
		url := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", config.Cfg.Region, config.Cfg.CognitoPoolId)
		options := keyfunc.Options{
			RefreshErrorHandler: func(err error) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			},
			RefreshInterval:   time.Hour,
			RefreshRateLimit:  time.Minute * 5,
			RefreshTimeout:    time.Second * 10,
			RefreshUnknownKID: true,
		}
		jwks, err := keyfunc.Get(url, options)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		claims, err := ValidateToken(tokenString, jwks)
		if err != nil {
			log.Println("error: ", err)
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "Invalid token",
			})
			c.Abort()
			return
		}
		c.Set("user-id", claims["username"].(string))
		switch v := claims["cognito:groups"].(type) {
		case []interface{}:
			for _, a := range v {
				array := a.(string)
				c.Set("cognito-group", array)
			}
		}
		cognitoGroup := c.GetString("cognito-group")
		if len(cognitoGroup) == 0 {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: "user doesn't belongs to any group",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ValidateToken(tokenString string, jwks *keyfunc.JWKS) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}

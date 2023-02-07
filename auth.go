package z_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Authorizer
// middleware for gin, abort the request if token is invalid
// config: Configuration created by NewConfiguration or DefaultConfiguration
// return: gin.Handler
func Authorizer(config Configuration) gin.HandlerFunc {
	return func(context *gin.Context) {
		// get the token from the request
		token := GetTokenFromContext(context, config)
		// if token is empty, return 401
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "token is required",
			})
			context.Abort()
			return
		}

		// parse the token
		claims, err := ParseToken(token, config)
		// if token is invalid, return 401
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "token is invalid",
			})
			context.Abort()
			return
		}

		// check the token if valid
		if CheckTokenValid(*claims) {
			context.Next()
		} else {
			// if token is out of date, return 401
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "token is out of date",
			})
			context.Abort()
			return
		}
	}
}

// GetTokenFromContext get the token from the request
// config: Configuration created by NewConfiguration or DefaultConfiguration
// return: token string
func GetTokenFromContext(c *gin.Context, config Configuration) string {
	// find the token from the request
	token := c.GetHeader(config.ParamName)
	if token == "" {
		token = c.Query(config.ParamName)
	}
	if token == "" {
		token = c.PostForm(config.ParamName)
	}
	if token == "" {
		token = c.Request.Header.Get(config.ParamName)
	}
	return token
}

// CheckTokenValid check the token if is out of date, return true if valid
// claims: Claims struct, get from ParseToken
// return: bool, true if valid
func CheckTokenValid(claims Claims) bool {
	// check the token if is out of date
	if time.Now().Unix() >= claims.ExpiresAt {
		return false
	} else {
		return true
	}
}

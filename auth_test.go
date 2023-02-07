package z_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	// create the configuration
	config := DefaultConfiguration()

	// create the router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// add the middleware
	router.Use(Authorizer(config))

	// add the test route
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "ok",
		})
	})

	// generate token
	token, _ := GenerateToken(config)
	fmt.Printf("token: \n%v\n", token)

	// request with token
	request, _ := http.NewRequest("GET", "/test", nil)
	request.Header.Set(config.ParamName, token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Printf("response with token: code: %v, body: %v\n", w.Code, w.Body.String())

	// request without token
	request, _ = http.NewRequest("GET", "/test", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Printf("response without token: code: %v, body: %v\n", w.Code, w.Body.String())

	// request with invalid token
	request, _ = http.NewRequest("GET", "/test", nil)
	request.Header.Set(config.ParamName, "invalid token")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Printf("response with invalid token: code: %v, body: %v\n", w.Code, w.Body.String())

	// request with expired token
	request, _ = http.NewRequest("GET", "/test", nil)
	expiredToken, _ := GenerateToken(NewConfiguration(config.JwtSecret, config.UserName, config.Signer, 0, config.ParamName))
	request.Header.Set(config.ParamName, expiredToken)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Printf("response with expired token: code: %v, body: %v\n", w.Code, w.Body.String())

}

// test generate token and parse token
func TestGenerateToken(t *testing.T) {
	config := DefaultConfiguration()

	token, err := GenerateToken(config)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("token: %v\n", token)

	claims, err := ParseToken(token, config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("claims: %v\n", claims)
}

package z_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	config := DefaultConfiguration()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(Authorizer(config))
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "ok",
		})
	})

	token, _ := GenerateToken(config)
	fmt.Printf("token: \n%v\n", token)

	// request with token
	request, _ := http.NewRequest("GET", "/test", nil)
	request.Header.Set(config.ParamName, token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Printf("response with token: %v\n", w.Body.String())

	// request without token
	request, _ = http.NewRequest("GET", "/test", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Printf("response without token: %v\n", w.Body.String())

	// request with invalid token
	request, _ = http.NewRequest("GET", "/test", nil)
	request.Header.Set(config.ParamName, "invalid token")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Printf("response with invalid token: %v\n", w.Body.String())

	// request with expired token
	request, _ = http.NewRequest("GET", "/test", nil)
	expiredToken, _ := GenerateToken(NewConfiguration(config.JwtSecret, config.UserName, config.Signer, 0, config.ParamName))
	request.Header.Set(config.ParamName, expiredToken)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Printf("response with expired token: %v\n", w.Body.String())

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

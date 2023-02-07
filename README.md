## Z_AUTH

### usage

```
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

// set the token
request.Header.Set(config.ParamName, token)

// send the request and print the response
w := httptest.NewRecorder()
router.ServeHTTP(w, request)
fmt.Printf("response with token: %v\n", w.Body.String())
```

there are four kinds of response:
>response with token: code: 200, body: {"msg":"ok"}  
response without token: code: 401, body: {"msg":"token is required"}  
response with invalid token: code: 401, body: {"msg":"token is invalid"}  
response with expired token: code: 401, body: {"msg":"token is out of date"}  
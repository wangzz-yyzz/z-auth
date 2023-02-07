package z_auth

type Configuration struct {
	JwtSecret  string
	UserName   string
	ExpireTime int
	Signer     string
	ParamName  string
}

// DefaultConfiguration returns a default configuration
// return: Configuration
func DefaultConfiguration() Configuration {
	return Configuration{
		JwtSecret:  "abc",
		UserName:   "admin",
		ExpireTime: 3,
		Signer:     "z-auth",
		ParamName:  "Authorization",
	}
}

// NewConfiguration returns a new configuration
// jwtSecret: jwt secret
// userName: username
// signer: jwt signer
// expireTime: jwt expire time (hour)
// paramName: jwt param name
// return: Configuration
func NewConfiguration(jwtSecret string, userName string,
	signer string, expireTime int, paramName string) Configuration {
	return Configuration{
		JwtSecret:  jwtSecret,
		UserName:   userName,
		ExpireTime: expireTime,
		Signer:     signer,
		ParamName:  paramName,
	}
}

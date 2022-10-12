package midlleware

// import (
// 	"errors"
// 	"hacktiv8-learning/final-project/models"
// 	"hacktiv8-learning/final-project/rules"
// 	"hacktiv8-learning/final-project/utils"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	jwt "github.com/appleboy/gin-jwt/v2"
// 	"github.com/gin-gonic/gin"
// )

// func InitUserMiddleware() *jwt.GinJWTMiddleware {
// 	// the jwt middleware
// 	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
// 		Realm:       "Register-Domain-user",
// 		Key:         []byte(os.Getenv("APP_KEY")),
// 		Timeout:     time.Hour * 24, //1 day
// 		MaxRefresh:  time.Hour * 24,
// 		IdentityKey: os.Getenv("APP_NAME"),
// 		PayloadFunc: func(data interface{}) jwt.MapClaims {
// 			if v, ok := data.(*models.User); ok {
// 				return map[string]interface{}{
// 					"id":      v.Id,
// 					"email":   v.Email,
// 					"role_id": v.RoleId,
// 				}
// 			}
// 			return jwt.MapClaims{}
// 		},
// 		IdentityHandler: func(c *gin.Context) interface{} {
// 			claims := jwt.ExtractClaims(c)
// 			return &models.User{
// 				Id:    uint64(claims["id"].(float64)),
// 				Email: claims["email"].(string),
// 			}
// 		},
// 		Authenticator: func(c *gin.Context) (interface{}, error) {
// 			var login rules.LoginForm

// 			if err := c.ShouldBindJSON(&login); err != nil {
// 				return "", jwt.ErrMissingLoginValues
// 			}
// 			//check email exist
// 			user, err := models.FindUserByEmail(login.Email)
// 			if err != nil {
// 				return nil, errors.New("email not found")
// 			}
// 			//user must active
// 			if user.IsActive == int64(0) {
// 				return nil, errors.New("account not active, please verify your email first")
// 			}
// 			err = utils.CompareHash(user.Password, login.Password)
// 			if err != nil {
// 				return nil, errors.New("email or password incorect")
// 			}
// 			return &user, nil
// 		},
// 		Authorizator: func(data interface{}, c *gin.Context) bool {
// 			if _, ok := data.(*models.User); ok {
// 				return true
// 			}

// 			return false
// 		},
// 		Unauthorized: func(c *gin.Context, code int, message string) {
// 			c.JSON(code, gin.H{
// 				"code":     code,
// 				"messages": message,
// 			})
// 		},
// 		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
// 			c.JSON(http.StatusOK, gin.H{
// 				"code":       http.StatusOK,
// 				"token":      token,
// 				"expire":     expire.Format(time.RFC3339),
// 				"token_type": "Bearer",
// 			})
// 		},
// 		SendCookie:     true,
// 		SecureCookie:   false,
// 		CookieHTTPOnly: true,
// 		CookieDomain:   os.Getenv("APP_FE"),
// 		CookieName:     "token",

// 		// TokenLookup is a string in the form of "<source>:<name>" that is used
// 		// to extract token from the request.
// 		// Optional. Default value "header:Authorization".
// 		// Possible values:
// 		// - "header:<name>"
// 		// - "query:<name>"
// 		// - "cookie:<name>"
// 		// - "param:<name>"
// 		TokenLookup: "header: Authorization, query: token, cookie: jwt",
// 		// TokenLookup: "query:token",
// 		// TokenLookup: "cookie:token",

// 		// TokenHeadName is a string in the header. Default value is "Bearer"
// 		TokenHeadName: "Bearer",

// 		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
// 		TimeFunc: time.Now,
// 	})

// 	if err != nil {
// 		log.Fatal("JWT Error:" + err.Error())
// 	}
// 	return authMiddleware
// }

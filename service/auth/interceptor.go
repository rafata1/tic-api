package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rafata1/tic-api/service/common"
)

type IService interface {
	AuthenticationInterceptor() gin.HandlerFunc
}

type service struct {
	iamClient IClient
}

// NewService ...
func NewService(iamEndpoint string) IService {
	return &service{iamClient: NewClient(iamEndpoint)}
}

const authorization = "authorization"
const bearerPrefix = "Bearer "
const emailKey = "email"

func (s service) AuthenticationInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorization)

		if len(header) < len(bearerPrefix) {
			common.WriteError(c, ErrInvalidIAMToken)
			c.Abort()
			return
		}

		if header[:len(bearerPrefix)] != bearerPrefix {
			common.WriteError(c, ErrInvalidIAMToken)
			c.Abort()
			return
		}

		token := header[len(bearerPrefix):]
		userInfo, err := s.iamClient.GetUserInfo(token)
		if err != nil {
			common.WriteError(c, err)
			c.Abort()
			return
		}

		c.Set(emailKey, userInfo.Email)
		c.Next()
	}
}

func GetUserEmail(c context.Context) string {
	email := c.Value(emailKey)
	return email.(string)
}

package mid

import (
	"strings"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/natholdallas/templates/fibergorm/pkg/conf"
	"github.com/natholdallas/templates/fibergorm/pkg/db"
	"github.com/natholdallas/templates/fibergorm/pkg/fibers"
)

type Claims struct {
	ID           uint
	Username     string
	IsAdmin      bool
	IsGoogleUser bool
	jwt.RegisteredClaims
}

var IsLogin = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(conf.Ctx.SecretKey)},
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return fibers.Err(err, fiber.StatusUnauthorized)
	},
})

func GenerateJwt(user db.User, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:           user.ID,
		Username:     user.Username,
		IsAdmin:      user.IsAdmin,
		IsGoogleUser: user.IsGoogleUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 120)},
		},
	})
	return token.SignedString([]byte(secretKey))
}

func ParseJwt(tokenstring, secretKey string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}

func FindClaims(c *fiber.Ctx) (Claims, error) {
	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	claims, err := ParseJwt(token, conf.Ctx.SecretKey)
	return *claims, err
}

func GetClaims(c *fiber.Ctx) Claims {
	claims, _ := FindClaims(c)
	return claims
}

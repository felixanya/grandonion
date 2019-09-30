package middleware

import (
    "errors"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "time"
)

func JwtAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.Request.Header.Get("token")
        if token == "" {
            c.JSON(http.StatusOK, gin.H{
                "status": -1,
                "msg": "请求未携带token, 无权限访问",
            })
            c.Abort()
            return
        }
        log.Println("---> get token:", token)

        j := NewJwt()
        claims, err := j.ParseToken(token)
        if err != nil {
            if err == TokenExpired {
                c.JSON(http.StatusOK, gin.H{
                    "status": -1,
                    "msg": "授权已过期",
                })
                c.Abort()
                return
            }
            c.JSON(http.StatusOK, gin.H{
                "status": -1,
                "msg": err.Error(),
            })
            c.Abort()
            return
        }
        c.Set("claims", claims)
    }
}

type Jwt struct {
    SigningKey  []byte
}

var (
    TokenExpired        error = errors.New("Token has been expired ")
    TokenNotValidYet    error = errors.New("Token not active yet ")
    TokenMalformed      error = errors.New("")
    TokenInvalid        error = errors.New("Invalid token ")
    SignKey             string = "asdfg"
)

type CustomClaims struct {
    ID      string  `json:"id"`
    Name    string  `json:"name"`
    Email   string  `json:"email"`

    jwt.StandardClaims
}

func NewJwt() *Jwt {
    return &Jwt{
        []byte(GetSignKey()),
    }
}

func GetSignKey() string {
    return SignKey
}

func SetSignKey(key string) string {
    SignKey = key
    return SignKey
}

func (j *Jwt) CreateToken(claim CustomClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
    return token.SignedString(j.SigningKey)
}

func (j *Jwt) ParseToken(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
        return j.SigningKey, nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors & jwt.ValidationErrorMalformed != 0 {
                return nil, TokenMalformed
            } else if ve.Errors & jwt.ValidationErrorExpired != 0 {
                return nil, TokenExpired
            } else if ve.Errors & jwt.ValidationErrorNotValidYet != 0 {
                return nil, TokenNotValidYet
            } else {
                return nil, TokenInvalid
            }
        }
    }
    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, TokenInvalid
}

// 更新token
func (j *Jwt) RefreshToken(tokenString string) (string, error) {
    jwt.TimeFunc = func() time.Time {
        return time.Unix(0, 0)
    }
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return j.SigningKey, nil
    })
    if err != nil {
        return "", err
    }
    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        jwt.TimeFunc = time.Now
        claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
        return j.CreateToken(*claims)
    }
    return "", TokenInvalid
}

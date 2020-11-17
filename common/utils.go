// Common tools and helper functions
package common

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// A helper function to generate random string
// 一个辅助函数，用于生成随机字符串
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 保持这两个配置私有，它们不应该是对外暴露的资源
const NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"
const NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"

// 一个工具函数，用于生成 jwt_token，可以在请求头中使用
func GenToken(id uint) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(NBSecretPassword))
	return token
}

// 自定义的 Error 类型，用于 API 接口返回自定义的错误信息
//  {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// 用于处理 gin 框架返回的 c.Bind 绑定错误（error）
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

// 封装错误信息到自定义的错误类型（CommonError）
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

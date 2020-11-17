package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/wangzitian0/golang-gin-starter-kit/common"
	"net/http"
	"strings"
)

// 从令牌中（Token）去除“TOKEN”前缀
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// 从授权头中提取令牌（Token）
// 使用 PostExtractionFilter 从标题中去除“TOKEN”前缀
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// 提取OAuth2访问令牌。  从 'Authorization' 请求头中查看
// 请求头中 'access_token' 的令牌参数
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// 一个辅助函数，获取用户模型并设置到 gin 框架上下文中
func UpdateContextUserModel(c *gin.Context, my_user_id uint) {
	var myUserModel UserModel

	// 如果传入的用户 id 值不为零就从数据库中查找该用户，并设置在框架上下文中
	if my_user_id != 0 {
		db := common.GetDB()
		db.First(&myUserModel, my_user_id)
	}
	c.Set("my_user_id", my_user_id)
	c.Set("my_user_model", myUserModel)
}

// 你能根据文档自定义中间件 https://github.com/gin-gonic/gin#custom-middleware
//  r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		// 从 http 请求（request）中提取 token
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.NBSecretPassword))
			return b, nil
		})
		// 如果解析 token 出错，不符合认证要求的如果设置了 auto401 就将http响应设置为返回 401 未认证返回
		// 否则就不做任何操作进入下一层处理
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}

		// 如果获取到了有效 token，就从 token 中解析出 token 中的数据设置认证用户模型到框架中
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			my_user_id := uint(claims["id"].(float64))
			//fmt.Println(my_user_id,claims["id"])
			UpdateContextUserModel(c, my_user_id)
		}
	}
}

package api

import (
	"demo-application-opa-golang/data"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful/v3"
	"time"
)

const SECRET = "fKAI2bjxgn2jqpx2"

func createToken(user data.User) (string,error) {
	var err error

	tokenClaims := jwt.MapClaims{}

	tokenClaims["user_id"] = user.UserId
	tokenClaims["user_name"] = user.Name
	tokenClaims["role"] = user.Role
	tokenClaims["exp"] = time.Now().Add(time.Minute * 4).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS384,tokenClaims)

	authToken, err := token.SignedString([]byte(SECRET))

	if err != nil {
		return "",err
	}

	return authToken, nil

}

func(api api) authFilter() restful.FilterFunction {
	return func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		store := api.Store
		authToken := request.Request.Header.Get("authorization")

		if authToken != ""  && store.SessionExists(authToken) {
			chain.ProcessFilter(request, response)
		} else {
			_ = response.WriteErrorString(401, "New Service Who Dis?")
		}
	}
}
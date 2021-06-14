package api

import (
	"context"
	"demo-application-opa-golang/data"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"io"
	"net/http"
)

type api struct {
	Store *data.Store
}

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func InitialiseWebService(store *data.Store) http.Handler {
	ws := new(restful.WebService)
	api := api{Store: store}
	restfulContainer := restful.NewContainer()

	ws.Filter(api.populateDataStoreInContext())
	ws.Route(ws.GET("/hello").Filter(api.authFilter()).To(api.hello))
	ws.Route(ws.POST("/login").To(api.login).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.DELETE("/logout").To(api.logout).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	restfulContainer.Add(ws)

	return restfulContainer.ServeMux
}

func (api api) hello(req *restful.Request, resp *restful.Response) {
	ctx := req.Request.Context()
	_ = resp.WriteAsJson(ctx.Value("store"))
}

func (api api) login(req *restful.Request, resp *restful.Response) {
	bytes, err := io.ReadAll(req.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
		_ = resp.WriteHeaderAndEntity(400, "Invalid Request")
	}
	loginRequestBody := loginRequest{}
	user := data.User{}

	err = json.Unmarshal(bytes, &loginRequestBody)

	if err != nil {
		_ = resp.WriteHeaderAndEntity(400, "Invalid Request")
	}

	user, err = api.Store.FindUserByName(loginRequestBody.Name)
	if err == nil {
		if user.Password != loginRequestBody.Password {
			_ = resp.WriteHeaderAndEntity(400, "Invalid Username or Password")
		}

		token, _ := createToken(user)

		loginResponseBody := loginResponse{Token: token}

		api.Store.Sessions = append(api.Store.Sessions, token)
		_ = resp.WriteAsJson(loginResponseBody)

	} else {
		_ = resp.WriteHeaderAndEntity(400, "Invalid Username or Password")
	}

}

func (api api) logout(req *restful.Request, resp *restful.Response) {
	authHeader := req.Request.Header.Get("Authorization")
	sessionToDelete := -10
	for index, session := range api.Store.Sessions {
		if session == authHeader {
			sessionToDelete = index
			break
		}
	}

	if sessionToDelete >= 0 {
		sessions := api.Store.Sessions
		sessions[sessionToDelete] = sessions[len(sessions)-1]
		api.Store.Sessions = sessions[:len(sessions)-1]
	}

	resp.WriteHeader(http.StatusOK)
}

func (api api) populateDataStoreInContext() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		r := req.Request
		ctx := context.WithValue(r.Context(), "store", api.Store)
		req.Request = r.WithContext(ctx)
		chain.ProcessFilter(req, resp)
	}
}

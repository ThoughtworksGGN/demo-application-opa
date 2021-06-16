package api

import (
	"context"
	"demo-application-opa-golang/data"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/rs/cors"
	uuid "github.com/satori/go.uuid"
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

type ticketsResponse struct {
	Tickets *[]data.Ticket `json:"tickets"`
}

type accountResponse struct {
	AccountId     uuid.UUID            `json:"account_id"`
	Payments      *[]data.Payment      `json:"payments"`
	Notifications *[]data.Notification `json:"notifications"`
	Tickets       *[]data.Ticket       `json:"tickets"`
	User          data.User            `json:"user"`
}

type accountForUserIdResponse struct {
	AccountId uuid.UUID `json:"account_id"`
}

type helloResponse struct {
	Store data.Store `json:"store"`
	RegoData data.RegoData `json:"rego_data"`
}

func InitialiseWebService(store *data.Store) http.Handler {
	ws := new(restful.WebService)
	api := api{Store: store}
	restfulContainer := restful.NewContainer()

	ws.Filter(api.populateDataStoreInContext())
	ws.Route(ws.GET("/hello").Filter(api.authFilter()).To(api.hello))
	ws.Route(ws.POST("/login").To(api.login).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.DELETE("/logout").To(api.logout).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/tickets").Filter(api.authFilter()).Filter(api.supportAuthzFilter()).To(api.getTickets).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/accounts/{accountId}").Filter(api.authFilter()).Filter(api.accountAuthzFilter()).To(api.getAccount).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/account").Filter(api.authFilter()).To(api.getAccountForUser).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	restfulContainer.Add(ws)

	corsRouter := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept"},
		AllowCredentials: true,
		Debug:            true,
		AllowedMethods:   []string{"HEAD", "GET", "POST", "DELETE"},
	})

	return corsRouter.Handler(restfulContainer.ServeMux)
}

func (api api) hello(_ *restful.Request, resp *restful.Response) {
	var helloResponse helloResponse
	regoData := api.Store.RegoData()
	helloResponse.Store = *api.Store
	helloResponse.RegoData = regoData

	_ = resp.WriteAsJson(helloResponse)
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

		api.Store.Sessions[token] = user
		_ = resp.WriteAsJson(loginResponseBody)

	} else {
		_ = resp.WriteHeaderAndEntity(400, "Invalid Username or Password")
	}

}

func (api api) logout(req *restful.Request, resp *restful.Response) {
	authHeader := req.Request.Header.Get("Authorization")

	delete(api.Store.Sessions, authHeader)

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

func (api api) getTickets(request *restful.Request, response *restful.Response) {
	tickets := api.Store.Tickets
	ticketsResponse := ticketsResponse{Tickets: &tickets}

	_ = response.WriteAsJson(ticketsResponse)
}

func (api api) getAccount(request *restful.Request, response *restful.Response) {
	accountId := uuid.FromStringOrNil(request.PathParameter("accountId"))
	accounts := api.Store.Accounts

	var matchingAccounts []data.Account

	for _, account := range accounts {
		if account.AccountId == accountId {
			matchingAccounts = append(matchingAccounts, account)
			break
		}
	}

	if len(matchingAccounts) == 1 {
		account := matchingAccounts[0]
		payments := api.Store.FindPaymentsByAccountId(account.AccountId)
		tickets := api.Store.FindTicketsByAccountId(account.AccountId)
		notifications := api.Store.FindNotificationsByAccountId(account.AccountId)
		user, _ := api.Store.FindUserByUserId(account.UserId)

		responseForAccount := accountResponse{
			AccountId:     account.AccountId,
			Payments:      &payments,
			Notifications: &notifications,
			Tickets:       &tickets,
			User:          user,
		}

		_ = response.WriteAsJson(responseForAccount)
	} else {
		_ = response.WriteErrorString(404, "Could not find matching account")
	}

}

func (api api) getAccountForUser(request *restful.Request, response *restful.Response) {
	sessionId := request.Request.Header.Get("authorization")
	user := api.Store.Sessions[sessionId]
	accounts := api.Store.Accounts

	var matchingAccounts []data.Account

	for _, account := range accounts {
		if account.UserId == user.UserId {
			matchingAccounts = append(matchingAccounts, account)
			break
		}
	}

	if len(matchingAccounts) == 1 {

		responseForAccount := accountForUserIdResponse{
			AccountId: matchingAccounts[0].AccountId,
		}
		_ = response.WriteAsJson(responseForAccount)
	} else {
		_ = response.WriteErrorString(404, "Could not find matching account")
	}

}

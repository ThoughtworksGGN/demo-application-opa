package api

import (
	"context"
	"demo-application-opa-golang/data"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
	"io/ioutil"
)

func isActionAllowed(input map[string]interface{}, queryText string, store *data.Store) bool {
	ctx := context.Background()
	var regoData map[string]interface{}
	regoDataAsJsonString, _ := json.Marshal(store.RegoData())

	json.Unmarshal(regoDataAsJsonString, &regoData)

	regoStore := inmem.NewFromObject(regoData)
	regoModule := regoModule()
	queryString := "data.api." + queryText

	query, _ := rego.New(rego.Query(queryString), rego.Store(regoStore), rego.Module("cerberus.rego", regoModule)).PrepareForEval(ctx)

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return false
	}
	if results[0].Expressions[0].Value == false {
		return false
	}
	return true
}

func (api api) accountAuthzFilter() restful.FilterFunction {
	return func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		query := "viewAccount"
		sessionId := request.Request.Header.Get("authorization")

		user := api.Store.Sessions[sessionId]
		input := make(map[string]interface{})
		input["userid"] = user.UserId
		input["accountid"] = request.PathParameter("accountId")

		if !isActionAllowed(input, query, api.Store) {
			_ = response.WriteErrorString(403, "Access Denied")
		} else {
			chain.ProcessFilter(request, response)
		}
	}
}

func (api api) supportAuthzFilter() restful.FilterFunction {
	return func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		query := "support"
		sessionId := request.Request.Header.Get("authorization")

		user := api.Store.Sessions[sessionId]
		input := make(map[string]interface{})
		input["userid"] = user.UserId

		if !isActionAllowed(input, query, api.Store) {
			_ = response.WriteErrorString(403, "Access Denied")
		} else {
			chain.ProcessFilter(request, response)
		}
	}
}

func regoModule() string {
	module := ""
	regoFile := "cerberus.rego"

	fileBytes, err := ioutil.ReadFile(regoFile)

	if err != nil {
		fmt.Println("Couldnot read rego file")
	}

	module = string(fileBytes)

	return module
}

package domain

import (
	"time"

	"github.com/federicoleon/golang-restclient/rest"
	//"github.com/mercadolibre/golang-restclient/rest" this version doesn't work with golang >= 1.13, so we are use the federicoleon while wait for approve the federicoleon PR
)

const (
	GrantTypePassowrd          = "password"
	GrantTypeClientCredentials = "client_credentials"
)

var (
	UserRestClient = rest.RequestBuilder{
		BaseURL: "micro_user", //microservice micro_user ip
		Timeout: 200 * time.Millisecond,
	}
)

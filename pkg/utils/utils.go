package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), err); err != nil {
			return
		}
	}
}

var PORT = "8080"
var POSTGRES_PORT = "5432"
var POSTGRES_USER = "postgres"
var POSTGRES_PASSWORD = "postgres"
var POSTGRES_DB = "postgres"
var POSTGRES_HOST = "localhost"

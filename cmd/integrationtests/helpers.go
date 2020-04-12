package integrationtests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

var router *httprouter.Router

func requestHelper(t *testing.T, method, path string, reader io.Reader, response interface{}) {
	request, err := http.NewRequest(method, path, reader)
	assert.Nil(t, err)

	dump, _ := httputil.DumpRequest(request, true)
	fmt.Println("--------------------------")
	fmt.Println("Request: ", strings.TrimSpace(string(dump)))

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, request)

	fmt.Println("Response: ", strings.TrimSpace(rr.Body.String()))
	fmt.Println("--------------------------")

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
}

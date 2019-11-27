package controllers

import (
	"encoding/json"
	"github.com/binatify/gin-template/example/app/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/atarantini/ginrequestid"
	"github.com/binatify/gin-template/base/context"
	"github.com/stretchr/testify/assert"
)

var (
	mockApp *Application
	mockCfg *AppConfig
)

func TestMain(m *testing.M) {
	mockApp = NewApplication("test", path.Clean("../.."))
	mockCfg = Config

	// inject middlewares and Resources
	{
		mockApp.Use("*", ginrequestid.RequestId(), context.NewLoggerMiddleware(mockApp.Logger()))
		mockApp.Resource()
	}

	// database clear
	code := m.Run()
	mongo := models.Model()
	mongo.Session().DB(mongo.Database()).DropDatabase()
	os.Exit(code)
}

func injectCookie(req *http.Request, cookie string) {
	req.Header.Set("Cookie", cookie)
}

func assertStatus(assertion *assert.Assertions, req *http.Request, expect int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	mockApp.ServeHTTP(w, req)
	assertion.Equal(expect, w.Code)
	return w
}

func assertStatusWithAuth(assertion *assert.Assertions, req *http.Request, cookie string, expect int) *httptest.ResponseRecorder {
	injectCookie(req, cookie)
	return assertStatus(assertion, req, expect)
}

func shouldBind(assertion *assert.Assertions, w *httptest.ResponseRecorder, item interface{}, isList ...bool) {
	body, err := ioutil.ReadAll(w.Body)
	assertion.Nil(err)
	assertion.Nil(json.Unmarshal(body, &item))
}

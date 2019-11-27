package ability

import (
	"encoding/json"
	"github.com/atarantini/ginrequestid"
	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/currentuser"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var (
	mockApp             *gin.Engine
	mockResourceLoader  *_mockResourceLoader
	mockResourceOwnerID = "mock_owner_id"
)

type mockResource struct {
	ID string `json:"id"`
}

func (r mockResource) GetOwnerID() string {
	return r.ID
}

type _mockResourceLoader struct{}

func (*_mockResourceLoader) LoadResource(id string) (Resource, error) {
	return mockResource{
		ID: mockResourceOwnerID,
	}, nil
}

func TestMain(m *testing.M) {
	mockApp = gin.Default()

	logger := logrus.New()
	logger.Out = os.Stdout

	mockApp.Use(ginrequestid.RequestId(), context.NewLoggerMiddleware(logger))

	mockApp.Use(func(ctx *gin.Context) {
		roleInt, _ := strconv.Atoi(ctx.GetHeader("role"))
		cUser := &currentuser.User{
			ID:   ctx.GetHeader("uid"),
			Role: currentuser.UserRole(roleInt),
		}

		ctx.Set(currentuser.SessionKey, cUser)
		ctx.Next()
	})

	// add resource checker
	AddHandle("resources.index", ShouldBeOne(currentuser.UserRole(1)))
	AddHandle("resources.show", NewResourceLoader(mockResourceLoader, []currentuser.UserRole{1}, func(user *currentuser.User, r Resource) bool {
		return r.GetOwnerID() == user.ID
	}))

	mockApp.Handle(http.MethodGet, "/v1/resources", context.NewHandler(NewHandler(func(ctx *context.Context) {})))

	mockApp.Handle(http.MethodGet, "/v1/resources/:id", context.NewHandler(NewHandler(func(ctx *context.Context) {
		item := LoadResource(ctx)
		ctx.JSON(http.StatusOK, item)
	})))

	os.Exit(m.Run())
}

func TestAbility(t *testing.T) {
	assertion := assert.New(t)

	// check resource index
	{
		req, _ := http.NewRequest(http.MethodGet, "/v1/resources", nil)
		assertStatus(assertion, req, http.StatusForbidden)

		req.Header.Set("role", "1")
		assertStatus(assertion, req, http.StatusOK)
	}

	// check resource show
	{
		req, _ := http.NewRequest(http.MethodGet, "/v1/resources/1", nil)
		assertStatus(assertion, req, http.StatusForbidden)

		req.Header.Set("role", "1")
		req.Header.Set("uid", mockResourceOwnerID)
		w := assertStatus(assertion, req, http.StatusOK)

		body, err := ioutil.ReadAll(w.Body)
		assertion.Nil(err)

		var resp mockResource
		assertion.Nil(json.Unmarshal(body, &resp))

		assertion.Equal(mockResourceOwnerID, resp.ID)
	}
}

func assertStatus(assertion *assert.Assertions, req *http.Request, expect int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	mockApp.ServeHTTP(w, req)

	assertion.Equal(expect, w.Code)

	return w
}

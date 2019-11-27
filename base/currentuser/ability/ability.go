package ability

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/currentuser"
	"github.com/binatify/gin-template/base/errors"
)

var (
	abilities   = map[string]HandleFunc{}
	resourceKey = "_loaded_resource"
	
	RoutePrefix = "/v1/"
)

type Resource interface {
	GetOwnerID() string
}

type ResourceLoader interface {
	LoadResource(string) (Resource, error)
}

type CheckFunc func(*currentuser.User, Resource) bool

type HandleFunc func(ctx *context.Context) (bool, *errors.Error)

func NewHandler(fn func(*context.Context)) func(ctx *context.Context) {
	return func(ctx *context.Context) {
		if ok, err := doCheck(ctx); !ok {
			requestID := ctx.GetString(context.RequestId)
			ctx.JSON(err.Code, errors.NewErrorResponse(requestID, *err))
			return
		}
		fn(ctx)
	}
}

func LoadResource(ctx *context.Context) interface{} {
	return ctx.MustGet(resourceKey)
}

func AddHandle(action string, handler HandleFunc) {
	abilities[action] = handler
}

func ShouldBeOne(roles ...currentuser.UserRole) func(ctx *context.Context) (bool, *errors.Error) {
	return func(ctx *context.Context) (bool, *errors.Error) {
		user := currentuser.NewUserFromContext(ctx)

		if !user.InRoles(roles) {
			ctx.Logger().Errorf("user with role:<%v> can't load this resource", user.Role)
			return false, &errors.AccessDenied
		}

		return true, nil
	}
}

func NewResourceLoader(rl ResourceLoader, roles []currentuser.UserRole, check CheckFunc) func(ctx *context.Context) (bool, *errors.Error) {
	return func(ctx *context.Context) (bool, *errors.Error) {
		user := currentuser.NewUserFromContext(ctx)

		if !user.InRoles(roles) {
			ctx.Logger().Errorf("user with role:<%v>  can't load this resource", user.Role)
			return false, &errors.AccessDenied
		}

		id := ctx.Param("id")
		item, err := rl.LoadResource(id)
		if err != nil {
			ctx.Logger().Errorf("load resource with error(%v): %v", id, err)
			return false, &errors.InvalidParameter
		}

		ctx.Set(resourceKey, item)

		if ok := check(user, item); !ok {
			return false, &errors.AccessDenied
		}

		return true, nil
	}
}

func doCheck(ctx *context.Context) (bool, *errors.Error) {
	path := strings.TrimLeft(ctx.Request.URL.Path, RoutePrefix)

	splits := strings.Split(path, "/")
	resourceName := splits[0]

	handleFunc, ok := abilities[fmt.Sprintf("%s.%s", resourceName, resolveAction(ctx))]
	if !ok {
		handleFunc, ok = abilities[resourceName]
	}

	if !ok {
		return true, nil
	}
	return handleFunc(ctx)
}

func resolveAction(ctx *context.Context) string {
	withID := ctx.Param("id") != ""

	switch ctx.Request.Method {
	case http.MethodGet:
		if withID {
			return "show"
		}

		return "index"
	case http.MethodPost:
		return "create"
	case http.MethodPut:
		return "update"
	case http.MethodDelete:
		return "destroy"
	default:
		return ""
	}
}

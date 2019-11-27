package currentuser

import (
	"encoding/gob"
	"github.com/binatify/gin-template/base/context"
)

func init() {
	gob.Register(&User{})
}

type User struct {
	ID        string            `json:"id"`
	Code      string            `json:"code"`
	Role      UserRole          `json:"role"`
	Name      string            `json:"name"`
	ExtraData map[string]string `json:"extra_data"`
}

func NewUser(id, code string, role UserRole) *User {
	return &User{
		ID:   id,
		Code: code,
		Role: role,
	}
}

func (u *User) InRoles(roles []UserRole) bool {
	for _, role := range roles {
		if u.Role == role {
			return true
		}
	}

	return false
}

func (u *User) GetExtraField(key string) string {
	if u.ExtraData == nil {
		return ""
	}

	return u.ExtraData[key]
}

func (u *User) SetExtraField(key, value string) {
	if u.ExtraData == nil {
		u.ExtraData = map[string]string{
			key: value,
		}

		return
	}

	u.ExtraData[key] = value
}

func NewUserFromContext(ctx *context.Context) (user *User) {
	return ctx.MustGet(SessionKey).(*User)
}
